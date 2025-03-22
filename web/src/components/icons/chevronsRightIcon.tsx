import { IconProps } from "@/models/icon.props";
import React from "react";

const ChevronsRightIcon: React.FC<IconProps> = ({
  className,
  width,
  height,
}) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={width || "18"}
    height={height || "18"}
    className={className}
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    strokeWidth="2"
    strokeLinecap="round"
    strokeLinejoin="round"
  >
    <path d="M13 17l5-5-5-5" />
    <path d="M6 17l5-5-5-5" />
  </svg>
);

export default ChevronsRightIcon;
