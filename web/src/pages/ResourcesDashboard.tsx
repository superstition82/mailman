import { Button } from "@mui/joy";
import { useEffect, useMemo, useState } from "react";
import { toast } from "react-hot-toast";
import { useResourceStore } from "../store/module/resource";
import Icon from "../components/common/Icon";
import ResourceItem from "../components/ResourceItem";

function ResourcesDashboard() {
  const resourceStore = useResourceStore();
  const resources = resourceStore.state.resources;
  const [selectedList, _] = useState<ResourceId[]>([]);

  useEffect(() => {
    resourceStore.fetchResourceList().catch((error) => {
      console.log(error);
      toast.error(error.message);
    });
  }, []);

  const handleDeleteSelectedBtnClick = () => {
    if (selectedList.length == 0) {
      toast.error("선택된 파일이 존재하지 않습니다.");
    } else {
      // const warningText = "정말 삭제하시겠습니까?";
      selectedList.map(async (resourceId: ResourceId) => {
        await resourceStore.deleteResourceById(resourceId);
      });
    }
  };

  const resourceList = useMemo(
    () =>
      resources.map((resource) => {
        return <ResourceItem resource={resource} />;
      }),
    [resources]
  );

  return (
    <section className="w-full max-w-3xl min-h-full flex flex-col justify-start items-center px-4 sm:px-2 sm:pt-4 pb-8 bg-zinc-100">
      <div className="w-full relative">
        <div className="w-full flex flex-col justify-start items-start px-4 py-3 rounded-xl bg-white">
          <div className="relative w-full flex flex-row justify-between items-center">
            <p className="flex flex-row justify-start items-center select-none rounded">
              <Icon.Paperclip className="w-5 h-auto mr-1" /> 리소스
            </p>
          </div>
          <div className="w-full flex flex-row justify-end items-center space-x-2 mt-3 z-1">
            {selectedList.length > 0 && (
              <Button
                onClick={() => handleDeleteSelectedBtnClick()}
                color="danger"
              >
                <Icon.Trash2 className="w-4 h-auto" />
              </Button>
            )}
          </div>
          <div className="w-full flex flex-col justify-start items-start mt-4 mb-6">
            <div className="flex flex-col justify-start items-start w-full">
              {resources.length === 0 ? (
                <p className="w-full text-center text-base my-6 mt-8">
                  업로드된 파일이 존재하지 않습니다.
                </p>
              ) : (
                resourceList
              )}
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}

export default ResourcesDashboard;
