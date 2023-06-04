import { useEffect } from "react";
import { NavLink } from "react-router-dom";
import { useLayoutStore } from "../store/module/layout";
import { resolution } from "../utils/layout";
import Icon from "./Icon";

export const Header: React.FC = () => {
  const layoutStore = useLayoutStore();
  const showHeader = layoutStore.state.showHeader;

  useEffect(() => {
    const handleWindowResize = () => {
      if (window.innerWidth < resolution.sm) {
        layoutStore.setHeaderStatus(false);
      } else {
        layoutStore.setHeaderStatus(true);
      }
    };
    window.addEventListener("resize", handleWindowResize);
    handleWindowResize();
  }, [location]);

  return (
    <div
      className={`fixed sm:sticky top-0 left-0 w-full sm:w-56 h-full shrink-0 pointer-events-none sm:pointer-events-auto z-10 ${
        showHeader && "pointer-events-auto"
      }`}
    >
      <div
        className={`fixed top-0 left-0 w-full h-full bg-black opacity-0 pointer-events-none transition-opacity duration-300 sm:!hidden ${
          showHeader && "opacity-60 pointer-events-auto"
        }`}
        onClick={() => layoutStore.setHeaderStatus(false)}
      ></div>
      <header
        className={`relative w-56 sm:w-full h-full max-h-screen overflow-auto hide-scrollbar flex flex-col justify-start items-start py-4 z-30 bg-zinc-100 sm:bg-transparent sm:shadow-none transition-all duration-300 -translate-x-full sm:translate-x-0 ${
          showHeader && "translate-x-0 shadow-2xl"
        }`}
      >
        <div className="w-full px-2 py-2 flex flex-col justify-start items-start shrink-0 space-y-2">
          <>
            <NavLink
              to="/"
              id="header-home"
              className={({ isActive }) =>
                `${
                  isActive && "bg-white shadow"
                } w-full px-4 pr-5 py-2 rounded-full flex flex-row items-center text-lg text-gray-800 hover:bg-white hover:shadow`
              }
            >
              <>
                <Icon.Home className="mr-3 w-6 h-auto opacity-70" /> 홈
              </>
            </NavLink>
            <NavLink
              to="/review"
              id="header-review"
              className={({ isActive }) =>
                `${
                  isActive && "bg-white shadow"
                } w-full px-4 pr-5 py-2 rounded-full flex flex-row items-center text-lg text-gray-800 hover:bg-white hover:shadow`
              }
            >
              <>
                <Icon.Calendar className="mr-3 w-6 h-auto opacity-70" /> 제목
              </>
            </NavLink>
            <NavLink
              to="/resources"
              id="header-resources"
              className={({ isActive }) =>
                `${
                  isActive && "bg-white shadow"
                } w-full px-4 pr-5 py-2 rounded-full flex flex-row items-center text-lg text-gray-800 hover:bg-white hover:shadow`
              }
            >
              <>
                <Icon.Paperclip className="mr-3 w-6 h-auto opacity-70" /> 리소스
              </>
            </NavLink>
          </>
          <NavLink
            to="/explore"
            id="header-explore"
            className={({ isActive }) =>
              `${
                isActive && "bg-white shadow"
              } w-full px-4 pr-5 py-2 rounded-full flex flex-row items-center text-lg text-gray-800 hover:bg-white hover:shadow`
            }
          >
            <>
              <Icon.Hash className="mr-3 w-6 h-auto opacity-70" /> 탐색
            </>
          </NavLink>
          <>
            <NavLink
              to="/archived"
              id="header-setting"
              className={({ isActive }) =>
                `${
                  isActive && "bg-white shadow"
                } w-full px-4 pr-5 py-2 rounded-full flex flex-row items-center text-lg text-gray-800 hover:bg-white hover:shadow`
              }
            >
              <>
                <Icon.Archive className="mr-3 w-6 h-auto opacity-70" /> 아카이브
              </>
            </NavLink>
            <NavLink
              to="/setting"
              id="header-setting"
              className={({ isActive }) =>
                `${
                  isActive && "bg-white shadow"
                } w-full px-4 pr-5 py-2 rounded-full flex flex-row items-center text-lg text-gray-800 hover:bg-white hover:shadow`
              }
            >
              <>
                <Icon.Settings className="mr-3 w-6 h-auto opacity-70" /> 설정
              </>
            </NavLink>
            <div className="pr-3 pl-1 w-full">
              <button
                className="mt-2 w-full py-3 rounded-full flex flex-row justify-center items-center bg-green-600 text-white hover:shadow hover:opacity-90"
                // onClick={() => showMemoEditorDialog()}
              >
                <Icon.Edit3 className="w-4 h-auto mr-1" /> New
              </button>
            </div>
          </>
        </div>
      </header>
    </div>
  );
};
