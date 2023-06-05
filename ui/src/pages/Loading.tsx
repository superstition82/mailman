import Icon from "../components/Icon";

function Loading() {
  return (
    <div className="flex flex-row justify-center items-center w-full h-full bg-zinc-100">
      <div className="w-80 max-w-full h-full py-4 flex flex-col justify-center items-center">
        <Icon.Loader className="animate-spin" />
      </div>
    </div>
  );
}

export default Loading;
