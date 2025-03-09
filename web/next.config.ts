import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  /* config options here */
  compilerOptions: {
    // ...
    include: [
      "types/next-env.d.ts",
      "**/*.ts",
      "**/*.tsx",
      ".next/types/**/*.ts",
      "types/**/*.ts",
      "**/*.ts",
      "**/*.tsx",
    ],
    // ...
  },
};

export default nextConfig;
