"use client";

import { SearchIcon, ListIcon } from "@/components/icons";
import { useActiveProject } from "@/providers/project";
import React from "react";

const ActivitiesPage: React.FC = () => {
  const { activeProject } = useActiveProject();

  return (
    <div className="flex-1 p-6 overflow-y-auto">
      <div className="flex items-center gap-2 mb-6">
        <span className="text-gray-600 text-sm">Projects</span>
        <span className="text-gray-400">/</span>
        <span className="text-gray-600 text-sm">{activeProject?.name}</span>
        <span className="text-gray-400">/</span>
        <span className="text-gray-800 text-sm font-medium">
          Logged Activities
        </span>
      </div>
      <div className="flex items-center justify-between mb-6">
        <h1 className="text-2xl font-semibold">Activity Log</h1>
      </div>
      <div className="flex gap-3 mb-6">
        <div className="flex items-center bg-white border border-gray-300 rounded-md py-2 px-3 flex-1 max-w-md">
          <SearchIcon className="text-gray-500 mr-2" />
          <input
            type="text"
            className="border-none outline-none flex-1 text-sm"
            placeholder="Search users..."
          />
        </div>

        <button className="flex items-center gap-2 py-2 px-3 rounded-md bg-white border border-gray-300 text-gray-700 font-medium text-sm cursor-pointer hover:bg-gray-50">
          <ListIcon className="w-4 h-4" />
          Filter
        </button>
      </div>
    </div>
  );
};

export default ActivitiesPage;
