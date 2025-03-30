"use client";
import { useActiveProject } from "@/providers/project";
import React from "react";

const SecurityDetailsPage: React.FC = () => {
  const { activeProject } = useActiveProject();
  return (
    <div className="flex-1 p-6 overflow-y-auto">
      <div className="flex items-center gap-2 mb-6">
        <span className="text-gray-600 text-sm">Projects</span>
        <span className="text-gray-400">/</span>
        <span className="text-gray-600 text-sm">{activeProject?.name}</span>
        <span className="text-gray-400">/</span>
        <span className="text-gray-800 text-sm font-medium">Security</span>
      </div>
      <div className="flex items-center justify-between mb-6">
        <h1 className="text-2xl font-semibold">Security Configuration</h1>
      </div>
      <div className="container mx-auto px-4 py-8">
        {/*  Security Dashboard Overview */}
        <div className="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
          <div className="bg-white p-6 rounded-lg shadow-sm border border-gray-200">
            <div className="flex items-center mb-4">
              <svg
                className="w-6 h-6 text-blue-500 mr-2"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M3 15a4 4 0 004 4h9a5 5 0 10-.1-9.999 5.002 5.002 0 10-9.78 2.096A4.001 4.001 0 003 15z"
                ></path>
              </svg>
              <h2 className="text-lg font-semibold">Firewall Status</h2>
            </div>
            <p className="text-green-600 font-medium">Active</p>
            <p className="text-sm text-gray-500 mt-2">
              Last updated: Today, 10:45 AM
            </p>
          </div>
          <div className="bg-white p-6 rounded-lg shadow-sm border border-gray-200">
            <div className="flex items-center mb-4">
              <svg
                className="w-6 h-6 text-blue-500 mr-2"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"
                ></path>
              </svg>
              <h2 className="text-lg font-semibold">Access Control</h2>
            </div>
            <p className="text-green-600 font-medium">8 Whitelisted IPs</p>
            <p className="text-sm text-gray-500 mt-2">2 IP ranges configured</p>
          </div>
          <div className="bg-white p-6 rounded-lg shadow-sm border border-gray-200">
            <div className="flex items-center mb-4">
              <svg
                className="w-6 h-6 text-blue-500 mr-2"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"
                ></path>
              </svg>
              <h2 className="text-lg font-semibold">Threat Protection</h2>
            </div>
            <p className="text-green-600 font-medium">No threats detected</p>
            <p className="text-sm text-gray-500 mt-2">
              12 threats blocked this month
            </p>
          </div>
        </div>

        {/* Firewall Configuration */}
        <div className="bg-white p-6 rounded-lg shadow-sm border border-gray-200 mb-8">
          <h2 className="text-xl font-semibold mb-4 flex items-center">
            <svg
              className="w-5 h-5 text-blue-500 mr-2"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M3 15a4 4 0 004 4h9a5 5 0 10-.1-9.999 5.002 5.002 0 10-9.78 2.096A4.001 4.001 0 003 15z"
              ></path>
            </svg>
            Firewall Configuration
          </h2>
          <div className="mb-4 flex items-center">
            <span className="mr-2">Firewall Status:</span>
            <label className="relative inline-flex items-center cursor-pointer">
              <input type="checkbox" defaultChecked className="sr-only peer" />
              <div className="w-11 h-6 bg-gray-200 rounded-full peer peer-focus:ring-4 peer-focus:ring-blue-300 peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-0.5 after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
            </label>
          </div>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <h3 className="font-medium mb-2">Rate Limiting</h3>
              <div className="mb-3">
                <label className="block text-sm font-medium text-gray-700 mb-1">
                  Maximum Requests per Minute
                </label>
                <input
                  type="number"
                  value="60"
                  className="w-full p-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500"
                />
              </div>
              <div className="mb-3">
                <label className="block text-sm font-medium text-gray-700 mb-1">
                  Block Duration (minutes)
                </label>
                <input
                  type="number"
                  defaultValue="15"
                  className="w-full p-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500"
                />
              </div>
            </div>
            <div>
              <h3 className="font-medium mb-2">Protocol Settings</h3>
              <div className="space-y-2">
                <label className="flex items-center">
                  <input
                    type="checkbox"
                    checked
                    className="rounded text-blue-600 focus:ring-blue-500 h-4 w-4"
                  />
                  <span className="ml-2">Enable HTTPS Only</span>
                </label>
                <label className="flex items-center">
                  <input
                    type="checkbox"
                    checked
                    className="rounded text-blue-600 focus:ring-blue-500 h-4 w-4"
                  />
                  <span className="ml-2">Block FTP Access</span>
                </label>
                <label className="flex items-center">
                  <input
                    type="checkbox"
                    checked
                    className="rounded text-blue-600 focus:ring-blue-500 h-4 w-4"
                  />
                  <span className="ml-2">Enable DDOS Protection</span>
                </label>
                <label className="flex items-center">
                  <input
                    type="checkbox"
                    checked
                    className="rounded text-blue-600 focus:ring-blue-500 h-4 w-4"
                  />
                  <span className="ml-2">Block Suspicious Activity</span>
                </label>
              </div>
            </div>
          </div>
        </div>

        {/* Whitelist Configuration */}
        <div className="bg-white p-6 rounded-lg shadow-sm border border-gray-200 mb-8">
          <h2 className="text-xl font-semibold mb-4 flex items-center">
            <svg
              className="w-5 h-5 text-blue-500 mr-2"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"
              ></path>
            </svg>
            IP Whitelist Configuration
          </h2>
          <div className="mb-4">
            <p className="text-sm text-gray-500 mb-2">
              Whitelisted IPs will bypass firewall restrictions and rate
              limiting.
            </p>
            <div className="flex mb-4">
              <input
                type="text"
                placeholder="Add IP address (e.g. 192.168.1.1)"
                className="flex-grow p-2 border border-gray-300 rounded-l-md focus:ring-blue-500 focus:border-blue-500"
              />
              <button className="bg-blue-500 text-white px-4 py-2 rounded-r-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2">
                Add IP
              </button>
            </div>
            <div className="overflow-x-auto">
              <table className="min-w-full divide-y divide-gray-200">
                <thead className="bg-gray-50">
                  <tr>
                    <th
                      scope="col"
                      className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                    >
                      IP Address
                    </th>
                    <th
                      scope="col"
                      className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                    >
                      Added By
                    </th>
                    <th
                      scope="col"
                      className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                    >
                      Date Added
                    </th>
                    <th
                      scope="col"
                      className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                    >
                      Notes
                    </th>
                    <th
                      scope="col"
                      className="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider"
                    >
                      Actions
                    </th>
                  </tr>
                </thead>
                <tbody className="bg-white divide-y divide-gray-200">
                  <tr>
                    <td className="px-6 py-4 whitespace-nowrap">
                      192.168.1.10
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap">Admin</td>
                    <td className="px-6 py-4 whitespace-nowrap">
                      Mar 28, 2025
                    </td>
                    <td className="px-6 py-4">Office main server</td>
                    <td className="px-6 py-4 text-right">
                      <button className="text-red-600 hover:text-red-900">
                        Delete
                      </button>
                    </td>
                  </tr>
                  <tr>
                    <td className="px-6 py-4 whitespace-nowrap">10.0.0.15</td>
                    <td className="px-6 py-4 whitespace-nowrap">Admin</td>
                    <td className="px-6 py-4 whitespace-nowrap">
                      Mar 27, 2025
                    </td>
                    <td className="px-6 py-4">Developer workstation</td>
                    <td className="px-6 py-4 text-right">
                      <button className="text-red-600 hover:text-red-900">
                        Delete
                      </button>
                    </td>
                  </tr>
                  <tr>
                    <td className="px-6 py-4 whitespace-nowrap">
                      203.0.113.0/24
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap">System</td>
                    <td className="px-6 py-4 whitespace-nowrap">
                      Mar 25, 2025
                    </td>
                    <td className="px-6 py-4">Office IP range</td>
                    <td className="px-6 py-4 text-right">
                      <button className="text-red-600 hover:text-red-900">
                        Delete
                      </button>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>

        {/* Authentication Settings */}
        <div className="bg-white p-6 rounded-lg shadow-sm border border-gray-200 mb-8">
          <h2 className="text-xl font-semibold mb-4 flex items-center">
            <svg
              className="w-5 h-5 text-blue-500 mr-2"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M15 7a2 2 0 012 2m4 0a6 6 0 01-7.743 5.743L11 17H9v2H7v2H4a1 1 0 01-1-1v-2.586a1 1 0 01.293-.707l5.964-5.964A6 6 0 1121 9z"
              ></path>
            </svg>
            Authentication Settings
          </h2>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <h3 className="font-medium mb-2">Password Policy</h3>
              <div className="space-y-2">
                <label className="flex items-center">
                  <input
                    type="checkbox"
                    checked
                    className="rounded text-blue-600 focus:ring-blue-500 h-4 w-4"
                  />
                  <span className="ml-2">Require Complex Passwords</span>
                </label>
                <label className="flex items-center">
                  <input
                    type="checkbox"
                    checked
                    className="rounded text-blue-600 focus:ring-blue-500 h-4 w-4"
                  />
                  <span className="ml-2">Password Expiry (90 days)</span>
                </label>
                <label className="flex items-center">
                  <input
                    type="checkbox"
                    checked
                    className="rounded text-blue-600 focus:ring-blue-500 h-4 w-4"
                  />
                  <span className="ml-2">Prevent Password Reuse</span>
                </label>
              </div>
            </div>
            <div>
              <h3 className="font-medium mb-2">Two-Factor Authentication</h3>
              <div className="space-y-2">
                <label className="flex items-center">
                  <input
                    type="checkbox"
                    checked
                    className="rounded text-blue-600 focus:ring-blue-500 h-4 w-4"
                  />
                  <span className="ml-2">Require 2FA for all users</span>
                </label>
                <label className="flex items-center">
                  <input
                    type="checkbox"
                    checked
                    className="rounded text-blue-600 focus:ring-blue-500 h-4 w-4"
                  />
                  <span className="ml-2">Allow SMS Authentication</span>
                </label>
                <label className="flex items-center">
                  <input
                    type="checkbox"
                    checked
                    className="rounded text-blue-600 focus:ring-blue-500 h-4 w-4"
                  />
                  <span className="ml-2">Allow Authenticator Apps</span>
                </label>
              </div>
            </div>
          </div>
        </div>

        {/* Security Log */}
        <div className="bg-white p-6 rounded-lg shadow-sm border border-gray-200">
          <h2 className="text-xl font-semibold mb-4 flex items-center">
            <svg
              className="w-5 h-5 text-blue-500 mr-2"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"
              ></path>
            </svg>
            Security Event Logs
          </h2>
          <div className="overflow-x-auto">
            <table className="min-w-full divide-y divide-gray-200">
              <thead className="bg-gray-50">
                <tr>
                  <th
                    scope="col"
                    className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                  >
                    Event
                  </th>
                  <th
                    scope="col"
                    className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                  >
                    IP Address
                  </th>
                  <th
                    scope="col"
                    className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                  >
                    User
                  </th>
                  <th
                    scope="col"
                    className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                  >
                    Date & Time
                  </th>
                  <th
                    scope="col"
                    className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                  >
                    Status
                  </th>
                </tr>
              </thead>
              <tbody className="bg-white divide-y divide-gray-200">
                <tr>
                  <td className="px-6 py-4 whitespace-nowrap">Login Attempt</td>
                  <td className="px-6 py-4 whitespace-nowrap">192.168.1.45</td>
                  <td className="px-6 py-4 whitespace-nowrap">
                    admin@example.com
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap">
                    Mar 30, 2025 09:45:12
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap">
                    <span className="px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-green-100 text-green-800">
                      Success
                    </span>
                  </td>
                </tr>
                <tr>
                  <td className="px-6 py-4 whitespace-nowrap">
                    Firewall Rule Updated
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap">192.168.1.10</td>
                  <td className="px-6 py-4 whitespace-nowrap">
                    admin@example.com
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap">
                    Mar 29, 2025 15:32:01
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap">
                    <span className="px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-green-100 text-green-800">
                      Success
                    </span>
                  </td>
                </tr>
                <tr>
                  <td className="px-6 py-4 whitespace-nowrap">Login Attempt</td>
                  <td className="px-6 py-4 whitespace-nowrap">45.123.45.67</td>
                  <td className="px-6 py-4 whitespace-nowrap">
                    user@example.com
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap">
                    Mar 29, 2025 10:15:30
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap">
                    <span className="px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-red-100 text-red-800">
                      Failed
                    </span>
                  </td>
                </tr>
                <tr>
                  <td className="px-6 py-4 whitespace-nowrap">IP Blocked</td>
                  <td className="px-6 py-4 whitespace-nowrap">103.54.123.87</td>
                  <td className="px-6 py-4 whitespace-nowrap">Unknown</td>
                  <td className="px-6 py-4 whitespace-nowrap">
                    Mar 28, 2025 23:12:45
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap">
                    <span className="px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-yellow-100 text-yellow-800">
                      Warning
                    </span>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
          <div className="mt-4 flex justify-end">
            <button className="text-blue-600 hover:text-blue-800">
              View All Logs
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default SecurityDetailsPage;
