const fetch = (function () {
    const __fetch__ = globalThis.__fetch__;
    delete globalThis.__fetch__;

    function Response({Status, StatusText, Headers, Body}) {
        this.status = Status;
        this.statusText = StatusText;
        this.headers = Headers;
        this.body = Body;
    }

    Response.prototype.text = function () {
        return this.body;
    }
    Response.prototype.json = function () {
        return JSON.parse(this.body);
    }

    return function () {
        if (arguments.length == 0) {
            throw new TypeError("Failed to execute 'fetch': 1 argument required, but only 0 present.");
        }
        let options = {
            method: "GET",
            headers: {
                "User-Agent": "Gopeed client",
            },
        }
        const request = arguments[0];
        if (typeof request == "string") {
            options.url = request;
        } else if (request instanceof URL) {
            options.url = request.href;
        } else {
            if (request.headers) {
                request.headers = {...options.headers, ...request.headers};
            }
            options = {...options, ...request};
        }
        if (options.body && typeof request != "string") {
            options.body = JSON.stringify(options.body);
        }

        if (arguments.length > 1) {
            const init = arguments[1];
            if (typeof init != "object") {
                throw new TypeError("Failed to execute 'fetch': The provided value is not of type 'RequestInit'.");
            }
            if (init.headers) {
                init.headers = {...options.headers, ...init.headers};
            }
            options = {...options, ...init};
        }
        return new Promise((resolve, reject) => {
            try {
                const resp = __fetch__(options.method, options.url, options.headers, options.body);
                resolve(new Response(resp));
            } catch (e) {
                reject(e);
            }
        });
    }
})();