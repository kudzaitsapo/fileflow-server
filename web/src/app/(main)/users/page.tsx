"use client";
import {
  DownloadIcon,
  ListIcon,
  PlusIcon,
  SearchIcon,
  TrashIcon,
} from "@/components/icons";
import TablePagination from "@/components/table/pagination";
import { ProjectUser } from "@/models/project";
import { useAxios } from "@/providers/axios";
import { useActiveProject } from "@/providers/project";
import { formatDateTime } from "@/utils/common";
import { useSession } from "next-auth/react";
import React, { useEffect, useState } from "react";

const UsersPage: React.FC = () => {
  const [users, setUsers] = useState<ProjectUser[]>([]);
  const { get } = useAxios();
  const { data: session } = useSession();
  const [total, setTotal] = useState<number>(0);
  const [currentPage, setCurrentPage] = useState<number>(1);
  const [pageSize, setPageSize] = useState<number>(10);
  const [isLoading, setIsLoading] = useState<boolean>(false);

  const offset = (currentPage - 1) * pageSize;

  const { activeProject } = useActiveProject();

  useEffect(() => {
    async function fetchUsers() {
      if (session && session.user && activeProject) {
        setIsLoading(true);
        const response = await get<ProjectUser[]>(
          `/projects/${activeProject.id}/users`,
          {
            limit: pageSize.toString(),
            offset: offset.toString(),
          }
        );
        if (response.success) {
          setIsLoading(false);
          setUsers(response.result);
          setTotal(response.meta.total_records);
        } else {
          setIsLoading(false);
          console.error("Failed to fetch users");
        }
      }
    }

    fetchUsers();
  }, [get, session, activeProject, pageSize, offset]);

  const handlePageChange = (page: number) => {
    setCurrentPage(page);
  };

  const handlePageSizeChange = (size: number) => {
    setPageSize(size);
  };

  return (
    <div className="flex-1 p-6 overflow-y-auto">
      <div className="flex items-center gap-2 mb-6">
        <span className="text-gray-600 text-sm">Projects</span>
        <span className="text-gray-400">/</span>
        <span className="text-gray-600 text-sm">{activeProject?.name}</span>
        <span className="text-gray-400">/</span>
        <span className="text-gray-800 text-sm font-medium">Users</span>
      </div>
      <div className="flex items-center justify-between mb-6">
        <h1 className="text-2xl font-semibold">Project Users</h1>
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

        <button className="flex items-center gap-2 py-2 px-3 rounded-md bg-blue-600 text-white border-none font-medium text-sm cursor-pointer hover:bg-blue-800">
          <PlusIcon className="w-4 h-4" />
          Create User
        </button>
      </div>
      <div className="bg-white rounded-lg border border-gray-200 overflow-hidden">
        <table className="w-full border-collapse">
          <thead>
            <tr>
              <th className="text-left py-3 px-4 text-xs uppercase text-gray-500 border-b border-gray-200 bg-gray-50">
                Name
              </th>
              <th className="text-left py-3 px-4 text-xs uppercase text-gray-500 border-b border-gray-200 bg-gray-50">
                Email
              </th>
              <th className="text-left py-3 px-4 text-xs uppercase text-gray-500 border-b border-gray-200 bg-gray-50">
                Date Created
              </th>
              <th className="text-left py-3 px-4 text-xs uppercase text-gray-500 border-b border-gray-200 bg-gray-50">
                Is Active?
              </th>
              <th className="text-left py-3 px-4 text-xs uppercase text-gray-500 border-b border-gray-200 bg-gray-50">
                Actions
              </th>
            </tr>
          </thead>
          <tbody>
            {!users.length && (
              <tr className="h-64">
                <td colSpan={5} className="text-center">
                  <div className="flex flex-col items-center justify-center p-8">
                    <div className="mb-4 text-gray-300">
                      <svg
                        xmlns="http://www.w3.org/2000/svg"
                        width="48"
                        height="48"
                        viewBox="0 0 24 24"
                        fill="none"
                        stroke="currentColor"
                        strokeWidth="1.5"
                        strokeLinecap="round"
                        strokeLinejoin="round"
                      >
                        <rect
                          x="3"
                          y="3"
                          width="18"
                          height="18"
                          rx="2"
                          ry="2"
                        ></rect>
                        <line x1="3" y1="9" x2="21" y2="9"></line>
                        <line x1="9" y1="21" x2="9" y2="9"></line>
                      </svg>
                    </div>
                    <h3 className="mb-2 text-xl text-gray-600">
                      No files found
                    </h3>
                    <p className="text-sm text-gray-500">
                      There are no files to display at this time.
                    </p>
                  </div>
                </td>
              </tr>
            )}

            {users &&
              users.length > 0 &&
              users.map((user) => (
                <tr key={user.id} className="hover:bg-gray-50">
                  <td className="py-3 px-4 border-b border-gray-200">
                    <div className="flex items-center gap-3">
                      {/* <div className="w-8 h-8 flex items-center justify-center rounded bg-red-50 text-red-500">
                        {/* <SVGDisplay svg={file.icon || ""} /> 
                      </div> */}
                      <div>
                        <div className="text-sm font-medium text-gray-800">
                          {user.user_info.first_name}&nbsp;&nbsp;
                          {user.user_info.last_name}
                        </div>
                      </div>
                    </div>
                  </td>
                  <td className="py-3 px-4 border-b border-gray-200 text-sm text-gray-600">
                    {user.user_info.email}
                  </td>
                  <td className="py-3 px-4 border-b border-gray-200 text-sm text-gray-600">
                    {formatDateTime(user.user_info.created_at)}
                  </td>
                  <td className="py-3 px-4 border-b border-gray-200 text-sm">
                    {user.user_info.is_active ? "Yes" : "No"}
                  </td>

                  <td className="py-3 px-4 border-b border-gray-200">
                    <div className="flex gap-1">
                      <button className="bg-transparent border-none cursor-pointer p-1.5 rounded text-gray-500 hover:bg-gray-100 hover:text-gray-900">
                        <DownloadIcon width={18} height={18} />
                      </button>
                      <button className="bg-transparent border-none cursor-pointer p-1.5 rounded text-gray-500 hover:bg-gray-100 hover:text-gray-900">
                        <TrashIcon width={18} height={18} />
                      </button>
                    </div>
                  </td>
                </tr>
              ))}
          </tbody>
        </table>

        {/* Add the pagination footer */}
        <TablePagination
          total={total}
          currentPage={currentPage}
          pageSize={pageSize}
          onPageChange={handlePageChange}
          onPageSizeChange={handlePageSizeChange}
          isLoading={isLoading}
        />
      </div>
    </div>
  );
};

export default UsersPage;
