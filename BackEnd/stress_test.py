import csv
import subprocess
import time
import requests
import pandas as pd
from io import StringIO


container_name = "daily-nutri-planner_backend"
image_name = "daily-nutri-planner_backend"
host = "localhost"
locust_port = "8089"
file_name = "combined_stats.csv"


def run_docker():
    print("Starting Docker container...")
    docker_process = subprocess.Popen(
        [
            "docker",
            "run",
            "-d",  # Run container in detached mode
            "--rm",
            "--name",
            container_name,
            "-p",
            "8080:8080",
            image_name,
        ],
        stdout=subprocess.PIPE,
        stderr=subprocess.PIPE,
    )
    return docker_process


def run_locust():
    print("Starting Locust...")
    locust_process = subprocess.Popen(
        [
            "locust",
            "-f",
            "locustfile.py",
            "--headless",
            "-u",
            "10",
            "-r",
            "10",
            "-H",
            f"http://{host}:{locust_port}",
        ],
        stdout=subprocess.PIPE,
        stderr=subprocess.PIPE,
    )
    return locust_process


def init_csv(file_name, locust_report):
    with open(file_name, "w", newline="") as csv_file:
        fieldnames = [
            "Time",
            "Container",
            "CPU %",
            "Mem Usage / Limit",
            "Mem %",
            "Net I/O",
            "Block I/O",
            "PIDs",
        ]
        fieldnames.extend(locust_report.keys())  # Add the locust column headers
        writer = csv.DictWriter(csv_file, fieldnames=fieldnames)
        writer.writeheader()


def get_docker_stats():
    docker_stats = subprocess.check_output(
        f"docker stats {container_name} --no-stream", shell=True
    )
    docker_stats = docker_stats.decode().split("\n")[1].split()
    return docker_stats


def get_locust_report():
    response = requests.get(f"http://{host}:{locust_port}/stats/requests/csv")
    locust_report = pd.read_csv(StringIO(response.text))
    locust_report = locust_report.iloc[-1].to_dict()
    return locust_report


def main():
    docker_process = run_docker()
    locust_process = run_locust()

    print("Waiting for processes to start generating data...")
    time.sleep(5)

    # Fetch the initial locust report to get the column headers
    locust_report = get_locust_report()
    init_csv(file_name, locust_report)

    try:
        for i in range(3):
            docker_stats = get_docker_stats()
            locust_report = get_locust_report()

            current_time = time.strftime("%Y-%m-%d %H:%M:%S", time.gmtime())
            csv_data = {
                "Time": current_time,
                "Container": docker_stats[1],
                "CPU %": docker_stats[2],
                "Mem Usage / Limit": docker_stats[3] + " / " + docker_stats[4],
                "Mem %": docker_stats[5],
                "Net I/O": docker_stats[6] + " / " + docker_stats[7],
                "Block I/O": docker_stats[8] + " / " + docker_stats[9],
                "PIDs": docker_stats[10],
                **locust_report,
            }

            # Write the data to the CSV file
            with open(file_name, "a", newline="") as csv_file:
                writer = csv.DictWriter(csv_file, fieldnames=csv_data.keys())
                writer.writerow(csv_data)
                csv_file.flush()  # Ensure that data is written to disk

            time.sleep(5)
            print(f"|====== record {i}th time ======|")
    finally:
        print("Terminating Docker and Locust processes...")
        locust_process.terminate()
        docker_process.terminate()


if __name__ == "__main__":
    main()
