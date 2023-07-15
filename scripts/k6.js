import http from "k6/http";
import { check, sleep } from "k6";

export let options = {
  stages: [
    { duration: "20s", target: 1000 },
    { duration: "10s", target: 1000 },
    { duration: "20s", target: 500 },
    { duration: "10s", target: 750 },
    { duration: "30s", target: 500 },
    { duration: "10s", target: 0 },
  ],
};

export default function () {
  // This will hit the endpoint http://pi_app_container:8080/api/v1/user/1
  // which is fixed endpoint for the user with id 1
  // you may change the id to any other id
  const response = http.get("http://pi_app_container:8080/api/v1/user/1", {
    headers: { Accepts: "application/json" },
  });
  check(response, { "status is 200": (r) => r.status === 200 });
  sleep(0.3);
}