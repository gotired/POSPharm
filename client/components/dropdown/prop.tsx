import { Dispatch, SetStateAction } from "react";

type itemDetail = {
  value: string;
  label: string;
};

export type DropdownProp = {
  value: string;
  onChange: (event: any) => void;
  defaultValue: string;
  items: itemDetail[];
};
