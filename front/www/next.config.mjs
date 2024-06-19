/** @type {import('next').NextConfig} */
const nextConfig = {
    async headers() {
        return [
            {
                // matching all API routes
                source: "/:path*",
                headers: [
                    { key: "Access-Control-Allow-Credentials", value: "true" },
                    { key: "Access-Control-Allow-Origin", value: "http://localhost:3000" }, // replace this your actual origin
                    { key: "Access-Control-Allow-Methods", value: "GET, DELETE, PATCH, POST, PUT, OPTIONS" },
                    { key: "Access-Control-Allow-Headers", value: "*" },
                ]
            }
        ]
    }
}

export default nextConfig;
