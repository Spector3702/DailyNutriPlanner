from locust import HttpUser, task, between

class QuickstartUser(HttpUser):
    wait_time = between(1, 2) # The user waits between 1 and 2 seconds between tasks

    @task
    def my_task(self):
        self.client.get("/api/v1/everydayNutri/") # Replace "/my-endpoint" with your endpoint
