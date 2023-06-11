import { useEffect } from "react";
import { NavLink } from "react-router-dom";
import { useLayoutStore } from "../../store/module/layout";
import { resolution } from "../../utils/layout";
import Icon from "./Icon";

function Header() {
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
              id="header-explore"
              className={({ isActive }) =>
                `${
                  isActive && "bg-white shadow"
                } w-full px-4 pr-5 py-2 rounded-2xl flex flex-row items-center text-lg text-gray-800 hover:bg-white hover:shadow`
              }
            >
              <>
                <Icon.Hash className="mr-3 w-6 h-auto opacity-70" /> 이메일
              </>
            </NavLink>
            <NavLink
              to="/sender"
              id="header-home"
              className={({ isActive }) =>
                `${
                  isActive && "bg-white shadow"
                } w-full px-4 pr-5 py-2 rounded-2xl flex flex-row items-center text-lg text-gray-800 hover:bg-white hover:shadow`
              }
            >
              <>
                <Icon.UserCircle className="mr-3 w-6 h-auto opacity-70" />
                발신자 관리
              </>
            </NavLink>
            <NavLink
              to="/recipient"
              id="header-review"
              className={({ isActive }) =>
                `${
                  isActive && "bg-white shadow"
                } w-full px-4 pr-5 py-2 rounded-2xl flex flex-row items-center text-lg text-gray-800 hover:bg-white hover:shadow`
              }
            >
              <>
                <Icon.UserPlus className="mr-3 w-6 h-auto opacity-70" />
                수신자 관리
              </>
            </NavLink>
            <NavLink
              to="/resource"
              id="header-resources"
              className={({ isActive }) =>
                `${
                  isActive && "bg-white shadow"
                } w-full px-4 pr-5 py-2 rounded-2xl flex flex-row items-center text-lg text-gray-800 hover:bg-white hover:shadow`
              }
            >
              <>
                <Icon.Paperclip className="mr-3 w-6 h-auto opacity-70" /> 리소스
              </>
            </NavLink>
            <div className="pr-3 pl-1 w-full">
              <NavLink
                to="/write"
                className="mt-2 w-full py-3 rounded-full flex flex-row justify-center items-center bg-green-600 font-medium text-white dark:opacity-80 hover:shadow hover:opacity-90"
              >
                <Icon.Edit3 className="w-4 h-auto mr-1" /> 템플릿 작성
              </NavLink>
            </div>
          </>
        </div>
      </header>
    </div>
  );
}

export default Header;
