import { useEffect, useState } from "react";
import { Checkbox, FormLabel, Input, Sheet, Table } from "@mui/joy";
import { TableVirtuoso } from "react-virtuoso";
import { useRecepientStore } from "../../store/module/recepient";
import Icon from "../Icon";

function RecepientManagementTable() {
  const recepientStore = useRecepientStore();
  const { recepients, selectedIds } = recepientStore.state;
  const [setting, setSetting] = useState({ waitTime: 2000 });

  useEffect(() => {
    recepientStore.fetchRecepients();
  }, []);

  const handleSelectAll = (event: React.ChangeEvent<HTMLInputElement>) => {
    recepientStore.toggleSelectAll(event.target.checked);
  };

  const handleClick = (id: number) => {
    recepientStore.toggleSelect(id);
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setSetting((prev) => ({
      ...prev,
      [e.target.name]: e.target.value,
    }));
  };

  return (
    <div className="flex flex-col rounded-sm px-2 mb-8 bg-white ">
      <Sheet
        sx={{
          "--TableCell-height": "40px",
          "--TableHeader-height": "calc(1 * var(--TableCell-height))",
        }}
      >
        <TableVirtuoso
          style={{ height: "75vh" }}
          data={recepients}
          components={{
            Table: (props) => (
              <Table
                stickyHeader
                hoverRow
                sx={{
                  "& thead th:nth-child(1)": {
                    width: "40px",
                  },
                  "& thead th:nth-child(2)": {
                    width: "80px",
                  },
                }}
                {...props}
              />
            ),
          }}
          fixedHeaderContent={() => (
            <tr>
              <th>
                <Checkbox
                  indeterminate={
                    selectedIds.length > 0 &&
                    selectedIds.length < recepients.length
                  }
                  checked={
                    recepients.length > 0 &&
                    selectedIds.length === recepients.length
                  }
                  onChange={handleSelectAll}
                  sx={{ verticalAlign: "sub" }}
                />
              </th>
              <th>#</th>
              <th>이메일</th>
              <th>전송가능</th>
            </tr>
          )}
          itemContent={(_, recepient) => (
            <>
              <td onClick={() => handleClick(recepient.id)}>
                <Checkbox
                  checked={selectedIds.includes(recepient.id)}
                  sx={{ verticalAlign: "top" }}
                />
              </td>
              <td>{recepient.id}</td>
              <td>{recepient.email}</td>
              <td>{recepient.reachable}</td>
            </>
          )}
        />
      </Sheet>
      <div className="p-2 flex justify-end gap-1">
        <>
          <FormLabel>대기시간(ms)</FormLabel>
          <Input
            className="w-20"
            size="sm"
            variant="soft"
            type="number"
            name="waitTime"
            value={setting.waitTime}
            onChange={handleChange}
            aria-label="초(ms)"
          />
          <button
            className="flex gap-2 px-4"
            onClick={() => {
              console.log("hello");
            }}
          >
            <Icon.MinusSquare className="w-6 h-auto opacity-70" /> 선택삭제
          </button>
        </>
      </div>
    </div>
  );
}

export default RecepientManagementTable;
