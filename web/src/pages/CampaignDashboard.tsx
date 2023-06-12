import { useCallback, useState } from "react";
import { Sheet, Table, Input, FormLabel } from "@mui/joy";
import { TableVirtuoso } from "react-virtuoso";
import Icon from "../components/common/Icon";
import { useLoading } from "../hooks/useLoading";

function CampaignDashboard() {
  const {} = useLoading();
  const [logs, setLogs] = useState([]);
  const [setting, setSetting] = useState({
    random: true,
    bundle: 1,
    waitTime: 2000,
  });

  const handleChangeInput = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      setSetting((prev) => ({
        ...prev,
        [e.target.name]: e.target.value,
      }));
    },
    [setSetting]
  );

  return (
    <section className="w-full max-w-3xl min-h-full flex flex-col justify-start items-center px-4 sm:px-2 sm:pt-4 pb-8 bg-zinc-100">
      <div className="w-full relative px-4 py-2 rounded-xl bg-white">
        <div className="relative w-full flex flex-row justify-between items-center">
          <p className="flex flex-row justify-start items-center select-none rounded">
            <Icon.AlertCircle className="w-5 h-auto mr-2" />
            템플릿, 발신자, 수신자 관리 탭에서 선택한 것이 반영되어요.
          </p>
        </div>
        <div className="w-full flex flex-row justify-end items-center space-x-2 mt-3 z-1">
          <div className="flex flex-col rounded-md px-2 mb-8 bg-white ">
            <Sheet
              sx={{
                "--TableCell-height": "40px",
                "--TableHeader-height": "calc(1 * var(--TableCell-height))",
              }}
            >
              <TableVirtuoso
                style={{ height: 400 }}
                data={[]}
                components={{
                  Table: (props) => (
                    <Table
                      stickyHeader
                      hoverRow
                      sx={{
                        "& thead th:nth-child(1)": {
                          width: "20%",
                        },
                      }}
                      {...props}
                    />
                  ),
                }}
                fixedHeaderContent={() => (
                  <tr>
                    <th>시간</th>
                    <th>제목</th>
                  </tr>
                )}
                itemContent={(_, log) => (
                  <>
                    <td>시간</td>
                    <td className="break-words">{log}</td>
                  </>
                )}
              />

              <div className="p-2 flex justify-end gap-1">
                <>
                  <FormLabel>랜덤발송:</FormLabel>
                  <input
                    className="mr-2 w-4"
                    type="checkbox"
                    name="random"
                    checked={setting.random}
                    onChange={({ target: { checked } }) => {
                      console.log(checked);
                    }}
                  />

                  <FormLabel>묶음발송: </FormLabel>
                  <Input
                    className="w-20"
                    size="sm"
                    variant="soft"
                    type="number"
                    name="bundle"
                    value={setting.bundle}
                    onChange={handleChangeInput}
                    aria-label="개"
                  />

                  <FormLabel>대기시간(ms): </FormLabel>
                  <Input
                    className="w-20"
                    size="sm"
                    variant="soft"
                    type="number"
                    name="waitTime"
                    value={setting.waitTime}
                    onChange={handleChangeInput}
                    aria-label="초(ms)"
                  />

                  <button
                    className="px-3 py-1 bg-blue-500 hover:bg-blue-700 text-white text-sm font-bold rounded"
                    onClick={() => {}}
                  >
                    메일 발송
                  </button>
                  <button
                    className="px-3 py-1 bg-blue-500 hover:bg-blue-700 text-white text-sm font-bold rounded"
                    onClick={() => {}}
                  >
                    숨은참조 발송
                  </button>
                </>
              </div>
            </Sheet>
          </div>
        </div>
      </div>
    </section>
  );
}

export default CampaignDashboard;
