"use client";
import React from "react";
import { useSession } from "next-auth/react";

const Header: React.FC = () => {
  const { data: session } = useSession();
  return (
    <>
      <header className="bg-white shadow sticky top-0 z-10 flex items-center justify-between px-6 py-3">
        <div className="flex items-center gap-2 font-bold text-xl text-blue-800">
          <span className="w-6 h-6 bg-blue-600 rounded flex items-center justify-center text-white font-bold">
            F
          </span>
          <span>FileFlow</span>
        </div>
        <div className="flex items-center gap-4">
          <div className="w-8 h-8 rounded-full bg-gray-300 flex items-center justify-center text-gray-700 font-semibold">
            {(session && session.user && session.user.name[0]) ?? `FF`}
          </div>
        </div>
      </header>
    </>
  );
};

export default Header;
