import { IconProps } from "@/models/icon.props";
import React from "react";

const EyeIcon: React.FC<IconProps> = ({ className, width, height }) => (
  <svg
    className={className}
    xmlns="http://www.w3.org/2000/svg"
    width={width || "20"}
    height={height || "20"}
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    strokeWidth="2"
    strokeLinecap="round"
    strokeLinejoin="round"
  >
    <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"></path>
    <circle cx="12" cy="12" r="3"></circle>
  </svg>
);

export default EyeIcon;
