import { ITableProps } from "@/models/table.props";
import React from "react";
import {
  ChevronLeftIcon,
  ChevronRightIcon,
  ChevronsLeftIcon,
  ChevronsRightIcon,
} from "../icons";

const TablePagination: React.FC<ITableProps> = ({
  total,
  currentPage,
  pageSize,
  onPageChange,
  onPageSizeChange,
  isLoading,
}) => {
  const totalPages = Math.ceil(total / pageSize);

  console.log("totalPages", totalPages);

  const getPageNumbers = () => {
    const pages = [];
    const maxVisiblePages = 5;

    if (totalPages <= maxVisiblePages) {
      // If we have fewer pages than our max visible, show them all
      for (let i = 1; i <= totalPages; i++) {
        pages.push(i);
      }
    } else {
      // Always show first page
      pages.push(1);

      // Calculate start and end of visible page range
      let start = Math.max(2, currentPage - 1);
      let end = Math.min(totalPages - 1, currentPage + 1);

      // Adjust for edge cases
      if (currentPage <= 2) {
        end = 4;
      } else if (currentPage >= totalPages - 2) {
        start = totalPages - 3;
      }

      // Add ellipsis if needed before middle pages
      if (start > 2) {
        pages.push("...");
      }

      // Add middle pages
      for (let i = start; i <= end; i++) {
        pages.push(i);
      }

      // Add ellipsis if needed after middle pages
      if (end < totalPages - 1) {
        pages.push("...");
      }

      // Always show last page
      if (totalPages > 1) {
        pages.push(totalPages);
      }
    }

    return pages;
  };

  const handlePageClick = (page: number | string) => {
    if (page === currentPage || page === "..." || isLoading) return;
    onPageChange(parseInt(page.toString()));
  };
  return (
    <div className="w-full px-4 py-3 flex items-center justify-between border-t border-gray-200 bg-white">
      <div className="flex items-center text-sm text-gray-500">
        <span>Showing </span>
        <select
          className="mx-1 rounded border border-gray-300 p-1 text-sm"
          value={pageSize}
          onChange={(e) => onPageSizeChange(Number(e.target.value))}
          disabled={isLoading}
        >
          <option value={10}>10</option>
          <option value={25}>25</option>
          <option value={50}>50</option>
          <option value={100}>100</option>
        </select>
        <span> of {total} items</span>
      </div>

      <div className="flex items-center space-x-1">
        <button
          onClick={() => handlePageClick(1)}
          disabled={currentPage === 1 || isLoading || totalPages === 0}
          className={`p-1 rounded ${
            currentPage === 1 || totalPages === 0
              ? "text-gray-300 cursor-not-allowed"
              : "text-gray-500 hover:bg-gray-100 hover:text-gray-900"
          }`}
          aria-label="First Page"
        >
          <ChevronsLeftIcon />
        </button>

        <button
          onClick={() => handlePageClick(currentPage - 1)}
          disabled={currentPage === 1 || isLoading || totalPages === 0}
          className={`p-1 rounded ${
            currentPage === 1 || totalPages === 0
              ? "text-gray-300 cursor-not-allowed"
              : "text-gray-500 hover:bg-gray-100 hover:text-gray-900"
          }`}
          aria-label="Previous Page"
        >
          <ChevronLeftIcon />
        </button>

        <div className="flex items-center space-x-1">
          {getPageNumbers().map((page, index) => (
            <button
              key={index}
              onClick={() => handlePageClick(page)}
              disabled={page === "..." || isLoading}
              className={`
                h-8 w-8 flex items-center justify-center rounded
                ${page === "..." ? "cursor-default" : "cursor-pointer"}
                ${
                  page === currentPage
                    ? "bg-blue-600 text-white"
                    : "hover:bg-gray-100"
                }
                ${page === "..." ? "text-gray-500" : "text-gray-700"}
              `}
            >
              {page}
            </button>
          ))}
        </div>

        <button
          onClick={() => handlePageClick(currentPage + 1)}
          disabled={currentPage === totalPages || isLoading || totalPages === 0}
          className={`p-1 rounded ${
            currentPage === totalPages || totalPages === 0
              ? "text-gray-300 cursor-not-allowed"
              : "text-gray-500 hover:bg-gray-100 hover:text-gray-900"
          }`}
          aria-label="Next Page"
        >
          <ChevronRightIcon />
        </button>

        <button
          onClick={() => handlePageClick(totalPages)}
          disabled={currentPage === totalPages || isLoading || totalPages === 0}
          className={`p-1 rounded ${
            currentPage === totalPages || totalPages === 0
              ? "text-gray-300 cursor-not-allowed"
              : "text-gray-500 hover:bg-gray-100 hover:text-gray-900"
          }`}
          aria-label="Last Page"
        >
          <ChevronsRightIcon />
        </button>
      </div>
    </div>
  );
};

export default TablePagination;
