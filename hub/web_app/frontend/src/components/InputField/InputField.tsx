import { ReactNode } from "react";
import styles from "./InputField.module.scss";

type InputFieldProps = {
  value: string;
  onChange: (value: string) => void;
  placeholder?: string;
  icon?: ReactNode;
  onIconClick?: () => void;
  error?: string;
  type?: "text";
};

export default function InputField({
  value,
  onChange,
  placeholder,
  icon,
  onIconClick,
  error,
  type = "text",
}: InputFieldProps) {
  return (
    <div className={styles.inputContainer}>
      <div className={`${styles.inputWrapper} ${error ? styles.error : ""}`}>
        <input
          type={type}
          className={styles.input}
          value={value}
          onChange={(e) => onChange(e.target.value)}
          placeholder={placeholder || "Enter the input..."}
        />
        {icon && (
          <div className={styles.icon} onClick={onIconClick}>
            {icon}
          </div>
        )}
      </div>
      {error && <div className={styles.errorMessage}>{error}</div>}
    </div>
  );
}
