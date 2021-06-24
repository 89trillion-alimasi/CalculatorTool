import locust
from locust import between, HttpUser, task


class QuickstartUser(HttpUser):
    wait_time = between(3, 4)

    @task
    def GetCal(self):
        self.client.get("/calculator?expr=1+2")
        self.client.get("/calculator?expr=3+5 / 2")
        self.client.get("/calculator?expr=3/2")



