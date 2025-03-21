"use client";

import React from "react";

interface SidebarItemProps {
  name: string;
  isActive: boolean;
  onClick?: () => void;
}

const SidebarItem: React.FC<SidebarItemProps> = ({
  name,
  isActive,
  onClick,
}) => {
  return (
    <div
      className={`flex items-center gap-3 p-2 rounded-md text-sm cursor-pointer ${
        isActive ? "bg-blue-500 text-white" : "text-gray-600 hover:bg-gray-100"
      }`}
      onClick={onClick}
    >
      <svg
        className={`w-4 h-4 ${isActive ? "text-white" : "text-gray-500"}`}
        xmlns="http://www.w3.org/2000/svg"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        strokeWidth="2"
        strokeLinecap="round"
        strokeLinejoin="round"
      >
        <path d="M22 19a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h5l2 3h9a2 2 0 0 1 2 2z"></path>
      </svg>
      {name}
    </div>
  );
};

export default SidebarItem;
