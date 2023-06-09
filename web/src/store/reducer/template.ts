import { PayloadAction, createSlice } from "@reduxjs/toolkit";
import { uniqBy } from "lodash-es";

type State = {
  templates: Template[];
  selectedIds: TemplateId[];
  isFetching: boolean;
};

const templateSlice = createSlice({
  name: "template",
  initialState: {
    templates: [],
    selectedIds: [],
    isFetching: true,
  } as State,
  reducers: {
    createTemplate: (state, action: PayloadAction<Template>) => {
      return {
        ...state,
        templates: state.templates.concat(action.payload),
      };
    },
    patchTemplate: (state, action: PayloadAction<Partial<Template>>) => {
      return {
        ...state,
        templates: state.templates.map((template) => {
          if (template.id === action.payload.id) {
            return {
              ...template,
              ...action.payload,
            };
          } else {
            return template;
          }
        }),
      };
    },
    deleteTemplate: (state, action: PayloadAction<TemplateId>) => {
      return {
        ...state,
        templates: state.templates.filter((template) => {
          return template.id !== action.payload;
        }),
      };
    },
    deleteBulkTemplate: (state, action: PayloadAction<TemplateId[]>) => {
      return {
        ...state,
        templates: state.templates.filter((template) => {
          return !action.payload.includes(template.id);
        }),
      };
    },
    upsertTemplates: (state, action: PayloadAction<Template[]>) => {
      return {
        ...state,
        templates: uniqBy([...state.templates, ...action.payload], "id"),
      };
    },
    setIsFetching: (state, action: PayloadAction<boolean>) => {
      return {
        ...state,
        isFetching: action.payload,
      };
    },
    toggleSelectAll: (state, action: PayloadAction<boolean>) => {
      return {
        ...state,
        selectedIds: action.payload
          ? state.templates.map((template) => template.id)
          : [],
      };
    },
    toggleSelect: (state, action: PayloadAction<number>) => {
      if (state.selectedIds.includes(action.payload)) {
        return {
          ...state,
          selectedIds: state.selectedIds.filter((id) => id !== action.payload),
        };
      } else {
        return {
          ...state,
          selectedIds: [...state.selectedIds, action.payload],
        };
      }
    },
  },
});

export const {
  createTemplate,
  patchTemplate,
  deleteBulkTemplate,
  deleteTemplate,
  upsertTemplates,
  setIsFetching,
  toggleSelect,
  toggleSelectAll,
} = templateSlice.actions;

export default templateSlice.reducer;
