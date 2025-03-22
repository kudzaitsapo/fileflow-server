import { IconProps } from "@/models/icon.props";
import React from "react";

const ChevronLeftIcon: React.FC<IconProps> = ({ className, width, height }) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={width || "18"}
    height={height || "18"}
    viewBox="0 0 24 24"
    className={className}
    fill="none"
    stroke="currentColor"
    strokeWidth="2"
    strokeLinecap="round"
    strokeLinejoin="round"
  >
    <path d="M15 18l-6-6 6-6" />
  </svg>
);

export default ChevronLeftIcon;
