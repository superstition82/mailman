import * as api from "../../helpers/api";
import store, { useAppSelector } from "../";
import { deleteResource, setResources } from "../reducer/resource";

export const useResourceStore = () => {
  const state = useAppSelector((state) => state.resource);

  return {
    state,
    async fetchResourceList(): Promise<Resource[]> {
      const { data } = (await api.getResourceList()).data;
      store.dispatch(setResources(data));
      return data;
    },
    async createResource(resourceCreate: ResourceCreate): Promise<Resource> {
      const { data } = (await api.createResource(resourceCreate)).data;
      const resourceList = state.resources;
      store.dispatch(setResources([data, ...resourceList]));
      return data;
    },
    async createResourceWithBlob(file: File): Promise<Resource> {
      const { name: filename } = file;
      const formData = new FormData();
      formData.append("file", file, filename);
      const { data } = (await api.createResourceWithBlob(formData)).data;
      const resourceList = state.resources;
      store.dispatch(setResources([data, ...resourceList]));
      return data;
    },
    async deleteResourceById(id: ResourceId) {
      await api.deleteResourceById(id);
      store.dispatch(deleteResource(id));
    },
  };
};
