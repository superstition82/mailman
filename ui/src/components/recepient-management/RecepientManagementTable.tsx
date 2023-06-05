import { useEffect } from "react";
import { Checkbox, Sheet, Table } from "@mui/joy";
import { TableVirtuoso } from "react-virtuoso";
import { useRecepientStore } from "../../store/module/recepient";
import Icon from "../Icon";

function RecepientManagementTable() {
  const recepientStore = useRecepientStore();
  const { recepients, selectedIds } = recepientStore.state;

  useEffect(() => {
    recepientStore.fetchRecepients();
  }, []);

  const handleSelectAll = (event: React.ChangeEvent<HTMLInputElement>) => {
    recepientStore.toggleSelectAll(event.target.checked);
  };

  const handleSelect = (id: number) => {
    recepientStore.toggleSelect(id);
  };

  const handleDeleteSelected = async () => {
    await recepientStore.deleteBulkRecepient(selectedIds);
  };

  const handleDeleteUnreachable = async () => {
    const unreableIds = recepients
      .filter((recepient) => recepient.reachable == "no")
      .map((recepient) => recepient.id);

    await recepientStore.deleteBulkRecepient(unreableIds);
  };

  const handleValidate = async () => {
    for (const id of selectedIds) {
      await recepientStore.validate(id);
    }
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
              <td onClick={() => handleSelect(recepient.id)}>
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
      <div className="py-4 flex justify-end gap-1">
        <>
          <button className="flex gap-2 px-4" onClick={handleValidate}>
            <Icon.CheckSquare className="w-6 h-auto opacity-70" /> 선택 검증
          </button>
          <button className="flex gap-2 px-4" onClick={handleDeleteSelected}>
            <Icon.MinusSquare className="w-6 h-auto opacity-70" /> 선택 삭제
          </button>
          <button className="flex gap-2 px-4" onClick={handleDeleteUnreachable}>
            <Icon.MinusSquare className="w-6 h-auto opacity-70" /> 전송불가 삭제
          </button>
        </>
      </div>
    </div>
  );
}

export default RecepientManagementTable;
