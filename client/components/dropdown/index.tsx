import * as React from "react";
import {
  Select as BaseSelect,
  SelectProps,
  SelectRootSlotProps,
} from "@mui/base/Select";
import {
  Option as BaseOption,
  OptionProps,
  OptionOwnerState,
} from "@mui/base/Option";
import { useTheme } from "@mui/system";
import clsx from "clsx";
import UnfoldMoreRoundedIcon from "@mui/icons-material/UnfoldMoreRounded";
import { MenuItem, SelectChangeEvent } from "@mui/material";
import { DropdownProp } from "./prop";

const getOptionColorClasses = ({
  selected,
  highlighted,
  disabled,
}: Partial<OptionOwnerState<number>>) => {
  let classes = "";
  if (disabled) {
    classes += " text-slate-400 dark:text-slate-700";
  } else {
    if (selected) {
      classes +=
        " bg-green-100 dark:bg-green-950 text-green-950 dark:text-green-50";
    } else if (highlighted) {
      classes +=
        " bg-slate-100 dark:bg-neutral-800 text-slate-900 dark:text-slate-300";
    }
    classes +=
      " hover:dark:bg-neutral-800 hover:bg-slate-100 hover:dark:text-slate-300 hover:text-slate-900";
    classes +=
      " focus-visible:outline focus-visible:outline-2 focus-visible:outline-green-400 focus-visible:dark:outline-green-300";
  }
  return classes;
};

const Option = React.forwardRef<HTMLLIElement, OptionProps<string>>(
  (props, ref) => {
    return (
      <BaseOption
        ref={ref}
        {...props}
        slotProps={{
          root: ({ selected, highlighted, disabled }) => ({
            className: `list-none p-2 rounded-lg cursor-default last-of-type:border-b-0 ${getOptionColorClasses(
              { selected, highlighted, disabled }
            )}`,
          }),
        }}
      />
    );
  }
);

const Button = React.forwardRef(function Button<
  TValue extends {},
  Multiple extends boolean
>(
  props: SelectRootSlotProps<TValue, Multiple>,
  ref: React.ForwardedRef<HTMLButtonElement>
) {
  const { ownerState, ...other } = props;
  return (
    <button type="button" {...other} ref={ref}>
      {other.children}
      <UnfoldMoreRoundedIcon className="m-auto" />
    </button>
  );
});

function useIsDarkMode() {
  const theme = useTheme();
  return theme.palette.mode === "dark";
}

export default function UnstyledSelectBasic({
  value,
  onChange,
  defaultValue,
  items,
}: DropdownProp) {
  const isDarkMode = useIsDarkMode();
  const handleChange = (value: string) => {
    onChange(value);
  };
  return (
    <div className={isDarkMode ? "dark" : ""}>
      <Select defaultValue={defaultValue} value={value}>
        {items.map((item, index) => (
          <Option
            key={index}
            onClick={() => {
              handleChange(item.value);
            }}
            value={item.value}
          >
            {item.label}
          </Option>
        ))}
      </Select>
    </div>
  );
}

const resolveSlotProps = (fn: any, args: any) =>
  typeof fn === "function" ? fn(args) : fn;

const Select = React.forwardRef(function CustomSelect<
  TValue extends {},
  Multiple extends boolean
>(
  props: SelectProps<TValue, Multiple>,
  ref: React.ForwardedRef<HTMLButtonElement>
) {
  // Replace this with your app logic for determining dark modes
  const isDarkMode = useIsDarkMode();

  return (
    <BaseSelect
      ref={ref}
      {...props}
      className={clsx("CustomSelect", props.className)}
      slots={{
        root: Button,
      }}
      slotProps={{
        ...props.slotProps,
        root: (ownerState) => {
          const resolvedSlotProps = resolveSlotProps(
            props.slotProps?.root,
            ownerState
          );
          return {
            ...resolvedSlotProps,
            className: clsx(
              `relative text-sm font-sans box-border w-75 px-3 py-2 rounded-lg text-left bg-white dark:bg-neutral-900 border border-solid border-slate-200 dark:border-neutral-700 text-slate-900 dark:text-neutral-300 transition-all hover:bg-slate-50 dark:hover:bg-neutral-800 outline-0 shadow-md shadow-slate-100 dark:shadow-slate-900 ${
                ownerState.focusVisible
                  ? "focus-visible:ring-4 ring-purple-500/30 focus-visible:border-purple-500 focus-visible:dark:border-purple-500"
                  : ""
              } [&>svg]:text-base [&>svg]:absolute [&>svg]:top-1/2 [&>svg]:-translate-y-1/2 [&>svg]:right-2.5`,
              resolvedSlotProps?.className
            ),
          };
        },
        listbox: (ownerState) => {
          const resolvedSlotProps = resolveSlotProps(
            props.slotProps?.listbox,
            ownerState
          );
          return {
            ...resolvedSlotProps,
            className: clsx(
              `text-sm font-sans p-1.5 my-3 w-75 rounded-xl overflow-auto outline-0 bg-white dark:bg-neutral-900 border border-solid border-slate-200 dark:border-neutral-700 text-slate-900 dark:text-neutral-300 shadow shadow-slate-200 dark:shadow-neutral-900`,
              resolvedSlotProps?.className
            ),
          };
        },
        popup: (ownerState) => {
          const resolvedSlotProps = resolveSlotProps(
            props.slotProps?.popup,
            ownerState
          );
          return {
            ...resolvedSlotProps,
            className: clsx(
              `${isDarkMode ? "dark" : ""} z-10`,
              resolvedSlotProps?.className
            ),
          };
        },
      }}
    />
  );
});
