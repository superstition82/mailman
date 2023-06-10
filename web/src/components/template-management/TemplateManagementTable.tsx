import { useEffect } from "react";
import { toast } from "react-hot-toast";
import { Button, Checkbox, Sheet, Table } from "@mui/joy";
import { TableVirtuoso } from "react-virtuoso";
import { useTemplateStore } from "../../store/module/template";
import { Link } from "react-router-dom";
import Icon from "../common/Icon";

function TemplateManagementTable() {
  const templateStore = useTemplateStore();
  const { templates, selectedIds } = templateStore.state;

  useEffect(() => {
    templateStore.fetchTemplates().catch((error) => {
      console.log(error);
      toast.error(error.message);
    });
  }, []);

  const handleSelectAll = (event: React.ChangeEvent<HTMLInputElement>) => {
    templateStore.toggleSelectAll(event.target.checked);
  };

  const handleClick = (_: React.MouseEvent, id: number) => {
    templateStore.toggleSelect(id);
  };

  const handleDeleteSelected = async () => {
    await templateStore.deleteBulkTemplate(selectedIds);
  };

  return (
    <div className="flex flex-col rounded-sm px-2 mb-8 bg-white">
      <Sheet
        sx={{
          "--TableCell-height": "40px",
          "--TableHeader-height": "calc(1 * var(--TableCell-height))",
        }}
      >
        <TableVirtuoso
          style={{ height: 400 }}
          data={templates}
          components={{
            Table: (props) => (
              <Table
                stickyHeader
                hoverRow
                sx={{
                  "& thead th:nth-child(1)": {
                    width: "40px",
                  },
                  "& thead th:nth-child(3)": {
                    width: "90px",
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
                    selectedIds.length < templates.length
                  }
                  checked={
                    templates.length > 0 &&
                    selectedIds.length === templates.length
                  }
                  onChange={handleSelectAll}
                  sx={{ verticalAlign: "sub" }}
                />
              </th>
              <th>제목</th>
              <th></th>
            </tr>
          )}
          itemContent={(_, template) => (
            <>
              <td onClick={(event) => handleClick(event, template.id)}>
                <Checkbox
                  checked={selectedIds.includes(template.id)}
                  sx={{ verticalAlign: "top" }}
                />
              </td>
              <td>{template.subject}</td>
              <td>
                <Link to={`/template/${template.id}`}>
                  <Button size="sm">수정하기</Button>
                </Link>
              </td>
            </>
          )}
        />
      </Sheet>
      <div className="py-4 flex justify-end gap-1">
        <button className="flex gap-2 px-4" onClick={handleDeleteSelected}>
          <Icon.MinusSquare className="w-6 h-auto opacity-70" /> 선택 삭제
        </button>
      </div>
    </div>
  );
}

export default TemplateManagementTable;
