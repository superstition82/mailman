import _ from "lodash-es";
import { useCallback, useState } from "react";
import { Sheet, Table, Input, FormLabel } from "@mui/joy";
import { TableVirtuoso } from "react-virtuoso";
import { toast } from "react-hot-toast";
import * as api from "../helpers/api";
import Icon from "../components/common/Icon";
import { useSenderStore } from "../store/module/sender";
import { useRecipientStore } from "../store/module/recipient";
import { useTemplateStore } from "../store/module/template";
import { splitArrayIntoGroups, wait } from "../helpers/f";
import { useLoading } from "../hooks/useLoading";
import Progress from "../components/common/Progress";

type Log = {
  content: string;
  createdAt: Date;
};

function CampaignDashboard() {
  const { progress, isLoading, setStart, setNextTick, setFinish } =
    useLoading();

  const senderStore = useSenderStore(),
    recipientStore = useRecipientStore(),
    templateStore = useTemplateStore();
  const { selectedIds: selectedSenderIds } = senderStore.state,
    { selectedIds: selectedRecipientIds } = recipientStore.state,
    { selectedIds: selectedTemplateIds } = templateStore.state;

  const [logs, setLogs] = useState<Log[]>([]);
  const [setting, setSetting] = useState({
    random: false,
    allBcc: false,
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

  const handleSendEmail = useCallback(async () => {
    // 1. 수신자 그룹을 만든다.
    // ex) ([1,2,3,4], 2) => [[1,2], [3,4]]
    const recipientIdGroups = _.chunk(selectedRecipientIds, setting.bundle);
    setStart(recipientIdGroups.length);

    // 2. 수신자 그룹을 발송자 개수로 나눈다.
    // ex) ([[1,2], [3,4]], 1) => [[[1,2], [3,4]]]
    const recipientIdGroupsBySender = splitArrayIntoGroups(
      recipientIdGroups,
      selectedSenderIds.length
    );

    try {
      // 3. 발송자 별로 묶은 수신자 그룹을 순회한다.
      for (let i = 0; i < selectedSenderIds.length; i++) {
        const recipientIdGroupBySender = recipientIdGroupsBySender[i];
        // 4. 랜덤이 아니면 템플릿을 순차적으로 발송한다.
        if (setting.random) {
          const templateId = _.sample(selectedTemplateIds) as number;
          for (const recipientIdGroup of recipientIdGroupBySender) {
            const result = await api.sendEmail({
              template: templateId,
              sender: selectedSenderIds[i],
              bcc: setting.allBcc ? recipientIdGroup : [],
              recipients: setting.allBcc ? [] : recipientIdGroup,
            });
            setNextTick();
            setLogs((prev) => [
              ...prev,
              {
                content: result.data,
                createdAt: new Date(),
              },
            ]);
            // 5. 정해진 시간만큼 기다린다.
            await wait(setting.waitTime);
          }
        } else {
          let templateIndex = 0;
          for (const recipientIdGroup of recipientIdGroupBySender) {
            const templateId = selectedTemplateIds[templateIndex++];
            const result = await api.sendEmail({
              template: templateId,
              sender: selectedSenderIds[i],
              bcc: setting.allBcc ? recipientIdGroup : [],
              recipients: setting.allBcc ? [] : recipientIdGroup,
            });
            setNextTick();
            setLogs((prev) => [
              ...prev,
              {
                content: result.data,
                createdAt: new Date(),
              },
            ]);
            if (templateIndex === selectedTemplateIds.length) {
              templateIndex = 0;
            }
            // 5. 정해진 시간만큼 기다린다.
            await wait(setting.waitTime);
          }
        }
      }
    } catch (error: unknown) {
      console.log(error);
      toast("이메일 전송 실패");
    } finally {
      setFinish();
    }
  }, [setting, selectedRecipientIds, selectedSenderIds, selectedTemplateIds]);

  return (
    <section className="w-full max-w-3xl min-h-full flex flex-col justify-start items-center px-4 sm:px-2 sm:pt-4 pb-8 bg-zinc-100">
      {isLoading ? <Progress value={progress} /> : <></>}
      <div className="w-full relative px-4 py-2 rounded-xl bg-white">
        <div className="relative w-full flex flex-row justify-between items-center">
          <p className="flex flex-row justify-start items-center select-none rounded">
            <Icon.Send className="w-5 h-auto mr-2" />
            이메일 발송
          </p>
        </div>
        <div className="w-full flex flex-row justify-end items-center space-x-2 mt-3 z-1">
          <div className="w-full flex flex-col rounded-md px-2 mb-8 bg-white ">
            <Sheet
              sx={{
                "--TableCell-height": "40px",
                "--TableHeader-height": "calc(1 * var(--TableCell-height))",
              }}
            >
              <TableVirtuoso
                style={{ height: 400 }}
                data={logs}
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
                    <td>{log.createdAt.toLocaleTimeString()}</td>
                    <td className="break-words">{log.content}</td>
                  </>
                )}
              />
              <div className="p-2 flex gap-1">
                <>
                  <FormLabel>랜덤발송:</FormLabel>
                  <input
                    className="mr-2 w-4 border"
                    type="checkbox"
                    name="random"
                    checked={setting.random}
                    onChange={({ target: { checked } }) => {
                      setSetting((prev) => ({ ...prev, random: checked }));
                    }}
                  />
                </>
                <>
                  <FormLabel>숨은참조:</FormLabel>
                  <input
                    className="mr-2 w-4 border"
                    type="checkbox"
                    name="random"
                    checked={setting.allBcc}
                    onChange={({ target: { checked } }) => {
                      setSetting((prev) => ({ ...prev, allBcc: checked }));
                    }}
                  />
                </>
                <>
                  <FormLabel>묶음발송: </FormLabel>
                  <Input
                    className="w-20 mr-2"
                    size="sm"
                    variant="soft"
                    type="number"
                    name="bundle"
                    value={setting.bundle}
                    onChange={handleChangeInput}
                    aria-label="개"
                  />
                </>
                <>
                  <FormLabel>간격(ms): </FormLabel>
                  <Input
                    className="w-20 mr-2"
                    size="sm"
                    variant="soft"
                    type="number"
                    name="waitTime"
                    value={setting.waitTime}
                    onChange={handleChangeInput}
                    aria-label="초(ms)"
                  />
                </>
                <div className="spacer flex-1" />
                <button
                  className="px-3 py-1 text-sm border rounded"
                  onClick={handleSendEmail}
                >
                  발송
                </button>
              </div>
            </Sheet>
          </div>
        </div>
      </div>
    </section>
  );
}

export default CampaignDashboard;
