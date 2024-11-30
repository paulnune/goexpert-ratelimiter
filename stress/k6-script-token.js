import http from 'k6/http';
import { htmlReport } from "https://raw.githubusercontent.com/benc-uk/k6-reporter/main/dist/bundle.js";

export default function () {
    let host = __ENV.RATELIMIT_HOST_TARGET
    let port = __ENV.RATELIMIT_PORT_TARGET
    let token_limit = __ENV.RATELIMIT_TOKEN_LIMIT_TARGET
    const URL = `http://${host}:${port}/hello`;
    const PARAMS = {
        headers: {
            'API_KEY': `Token${token_limit}`
        },
    };
    http.get(URL, PARAMS);
}

export function handleSummary(data) {
    return {
        "/home/k6/stress/summary-token.html": htmlReport(data),
    };
}
