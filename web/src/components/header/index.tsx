import React from "react";
import "@/app/globals.css";

const Header: React.FC = () => {
  return (
    <>
      <header>
        <div className="logo">
          <span className="logo-icon">F</span>
          <span>FileFlow</span>
        </div>
        <div className="user-menu">
          <div className="user-avatar">JS</div>
        </div>
      </header>
    </>
  );
};

export default Header;
