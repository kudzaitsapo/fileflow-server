import React from "react";

const SVGDisplay: React.FC<{ svg: string }> = ({ svg }) => {
  return (
    <div
      dangerouslySetInnerHTML={{
        __html: svg,
      }}
    />
  );
};

export default SVGDisplay;
