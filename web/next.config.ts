import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  // The floating dev badge sits on top of the bottom tab bar on a 360px
  // screen, which is exactly where the app is tested.
  devIndicators: false,
};

export default nextConfig;
