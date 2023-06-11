import ResourceItemDropdown from "./ResourceItemDropdown";

type Props = {
  resource: Resource;
};

function ResourceItem({ resource }: Props) {
  return (
    <div key={resource.id} className="px-2 py-2 w-full grid grid-cols-10">
      <span className="col-span-2 w-full m-auto truncate text-base pr-2">
        {resource.id}
      </span>
      <span className="col-span-6 w-full m-auto truncate text-base pr-2 hover:opacity-30">
        <a
          target="_blank"
          rel="noopener noreferrer"
          href={`/o/r/${resource.id}/${resource.filename}`}
        >
          {resource.filename} (
          {new Date(resource.createdTs * 1000).toLocaleTimeString()})
        </a>
      </span>
      <div className="col-span-1 w-full flex flex-row justify-end items-center pr-2">
        <ResourceItemDropdown resource={resource} />
      </div>
    </div>
  );
}

export default ResourceItem;
