type Props = {
  value: number;
};

function Progress({ value }: Props) {
  return (
    <div
      className="fixed top-0 left-0 right-0 bottom-0 flex items-center justify-center z-50"
      style={{ backgroundColor: "rgba(0,0,0,0.5)" }}
    >
      <div className="w-80 max-w-full h-full py-4 flex flex-col justify-center items-center text-gray-200">
        <div className="mt-2">{value.toFixed(2)}%</div>
        <div>(주의: 새로고침 시 작업이 중지됩니다.)</div>
      </div>
    </div>
  );
}

export default Progress;
