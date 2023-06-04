import React from "react";

export const Home: React.FC = () => {
  return (
    <div className="w-full flex flex-row justify-start items-start">
      <div className="flex-grow shrink w-auto px-4 sm:px-2 sm:pt-4">
        <div className="w-full h-auto flex flex-col justify-start items-start bg-zinc-100 rounded-lg"></div>
        {/* <EmailList/> */}
      </div>
    </div>
  );
};
