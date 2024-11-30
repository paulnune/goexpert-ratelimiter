import http from 'k6/http';
import { htmlReport } from "https://raw.githubusercontent.com/benc-uk/k6-reporter/main/dist/bundle.js";

export default function () {
    let host = __ENV.RATELIMIT_HOST_TARGET
    let port = __ENV.RATELIMIT_PORT_TARGET
    http.get(`http://${host}:${port}/hello`);
}

export function handleSummary(data) {
    return {
        "/home/k6/stress/summary-ip.html": htmlReport(data),
    };
}
