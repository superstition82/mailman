import copy from "copy-to-clipboard";
import toast from "react-hot-toast";
import Dropdown from "./common/Dropdown";
import Icon from "./common/Icon";
import { useResourceStore } from "../store/module/resource";
import { getResourceUrl } from "../utils/resource";

type Props = {
  resource: Resource;
};

function ResourceItemDropdown({ resource }: Props) {
  const resourceStore = useResourceStore();

  const handleCopyResourceLinkBtnClick = (resource: Resource) => {
    const url = getResourceUrl(resource);
    copy(url);
    toast.success("클립보드에 복사되었습니다");
  };

  const handleDeleteResourceBtnClick = async (resource: Resource) => {
    await resourceStore.deleteResourceById(resource.id);
  };

  return (
    <Dropdown
      actionsClassName="!w-auto min-w-[8rem]"
      trigger={
        <Icon.MoreVertical className="w-4 h-auto hover:opacity-80 cursor-pointer" />
      }
      actions={
        <>
          <button
            className="w-full text-left text-sm leading-6 py-1 px-3 cursor-pointer rounded hover:bg-gray-100"
            onClick={() => handleCopyResourceLinkBtnClick(resource)}
          >
            링크 복사
          </button>
          <button
            className="w-full text-left text-sm leading-6 py-1 px-3 cursor-pointer rounded text-red-600 hover:bg-gray-100"
            onClick={() => handleDeleteResourceBtnClick(resource)}
          >
            삭제
          </button>
        </>
      }
    />
  );
}

export default ResourceItemDropdown;
