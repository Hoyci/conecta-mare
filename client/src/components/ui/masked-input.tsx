import React from "react";
import InputMask from "react-input-mask";
import { Input } from "@/components/ui/input";

interface MaskedInputProps {
  mask?: string;
  alwaysShowMask?: boolean;
  value: string;
  onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
  onBlur?: (e: React.FocusEvent<HTMLInputElement>) => void;
  placeholder?: string;
  id?: string;
  className?: string;
}

export const MaskedInput = ({
  mask,
  alwaysShowMask = false,
  value,
  onChange,
  onBlur,
  placeholder,
  id,
  className,
}: MaskedInputProps) => {
  return (
    <InputMask
      mask={mask}
      alwaysShowMask={alwaysShowMask}
      maskChar={null}
      value={value}
      onChange={onChange}
      onBlur={onBlur}
    >
      {(inputProps: any) => (
        <Input
          {...inputProps}
          id={id}
          placeholder={placeholder}
          className={className}
        />
      )}
    </InputMask>
  );
};
