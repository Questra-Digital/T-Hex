import styles from "./LabelsInput.module.scss";
import {
  useState,
  useRef,
  useEffect,
  forwardRef,
  useImperativeHandle,
} from "react";
import { motion, AnimatePresence } from "framer-motion";

interface LabelsInputProps {
  initialLabels?: string[];
}

export interface LabelsInputRef {
  validate(): boolean;
  getData(): { labels: string[] };
  clearErrors(): void;
}

const LabelsInput = forwardRef<LabelsInputRef, LabelsInputProps>(
  ({ initialLabels = [] }, ref) => {
    const [labels, setLabels] = useState<string[]>(initialLabels);
    const [labelInput, setLabelInput] = useState("");
    const [showDropdown, setShowDropdown] = useState(false);
    const [inputWidth, setInputWidth] = useState(0);
    const [error, setError] = useState<string>("");
    const [showError, setShowError] = useState(false);

    const dropdownRef = useRef<HTMLDivElement>(null);
    const hiddenSpanRef = useRef<HTMLSpanElement>(null);
    const containerRef = useRef<HTMLDivElement>(null);

    // Validation function
    const validateLabels = (labelList: string[]): string => {
      if (labelList.length === 0) return "At least one label is required";
      return "";
    };

    // Expose methods to parent via ref
    useImperativeHandle(ref, () => ({
      validate(): boolean {
        const errorMessage = validateLabels(labels);
        setError(errorMessage);
        setShowError(!!errorMessage);
        return !errorMessage;
      },
      getData(): { labels: string[] } {
        return { labels };
      },
      clearErrors(): void {
        setError("");
        setShowError(false);
      },
    }));

    // Dynamically adjust input width
    useEffect(() => {
      if (hiddenSpanRef.current && labelInput.trim()) {
        const textWidth = hiddenSpanRef.current.offsetWidth;
        const paddingAndBorder = 30;
        const minWidth = 120;
        const maxWidth = 300;
        const calculatedWidth = Math.min(
          Math.max(textWidth + paddingAndBorder, minWidth),
          maxWidth
        );
        setInputWidth(calculatedWidth);
      }
    }, [labelInput]);

    // Example available labels
    const [availableLabels, setAvailableLabels] = useState([
      "Frontend",
      "Backend",
      "Bug",
      "Feature",
      "Urgent",
      "Blocked",
      "In Progress",
      "To Do",
      "Done",
      "In Review",
    ]);

    // Filtered labels
    const filteredLabels = availableLabels.filter(
      (label) =>
        label.toLowerCase().includes(labelInput.toLowerCase()) &&
        !labels.includes(label)
    );

    // ✅ Close dropdown when clicking outside
    useEffect(() => {
      const handleClickOutside = (event: MouseEvent) => {
        if (
          containerRef.current &&
          !containerRef.current.contains(event.target as Node)
        ) {
          setShowDropdown(false);
        }
      };

      document.addEventListener("mousedown", handleClickOutside);
      return () => {
        document.removeEventListener("mousedown", handleClickOutside);
      };
    }, []);

    // Add a label
    const handleAddLabel = (newLabel: string) => {
      const trimmed = newLabel.trim();
      if (!trimmed) return;

      if (!labels.some((l) => l.toLowerCase() === trimmed.toLowerCase())) {
        setLabels((prev) => [...prev, trimmed]);
      }

      if (
        !availableLabels.some(
          (l) => l.toLowerCase() === trimmed.toLowerCase()
        )
      ) {
        setAvailableLabels((prev) => [...prev, trimmed]);
      }

      setLabelInput("");
      setShowDropdown(true);

      // ✅ Clear error when label added
      setError("");
      setShowError(false);
    };

    const handleRemoveLabel = (labelToRemove: string) => {
      setLabels((prev) => prev.filter((l) => l !== labelToRemove));
    };

    return (
      <div
        ref={containerRef}
        className={`${styles.labelsInputContainer} ${error ? styles.error : ""}`}
        onClick={() => setShowDropdown(true)}
      >
        {/* Selected labels */}
        {labels.map((label, i) => (
          <span key={label + i} className={styles.labelChip}>
            {label}
            <button
              onClick={(e) => {
                e.stopPropagation();
                handleRemoveLabel(label);
              }}
              className={styles.removeLabel}
              type="button"
            >
              ×
            </button>
          </span>
        ))}

        {/* Input */}
        <input
          type="text"
          className={`${styles.labelInputField} ${
            labelInput.trim() ? styles.typingInChip : ""
          }`}
          value={labelInput}
          onChange={(e) => {
            setLabelInput(e.target.value);
            setShowDropdown(true);
          }}
          onFocus={() => {
            setShowDropdown(true);
            setError("");
            setShowError(false);
          }}
          onClick={() => setShowDropdown(true)}
          style={
            labelInput.trim()
              ? {
                  width: `${Math.max(inputWidth, 120)}px`,
                }
              : {}
          }
          placeholder={
            labels.length === 0
              ? "Select a label..."
              : "Add another label..."
          }
          onKeyDown={(e) => {
            if (e.key === "Enter" && labelInput.trim() !== "") {
              e.preventDefault();
              handleAddLabel(labelInput);
            }
          }}
        />

        {/* Hidden span for measuring width */}
        <span
          ref={hiddenSpanRef}
          style={{
            position: "absolute",
            visibility: "hidden",
            whiteSpace: "nowrap",
            fontSize: "16px",
            fontWeight: "500",
            fontFamily: "inherit",
            pointerEvents: "none",
          }}
        >
          {labelInput ||
            (labels.length === 0
              ? "Select a label..."
              : "Add another label...")}
        </span>

        {/* Dropdown */}
        <AnimatePresence>
          {showDropdown && (
            <motion.div
              ref={dropdownRef}
              initial={{ opacity: 0, y: -8 }}
              animate={{ opacity: 1, y: 0 }}
              exit={{ opacity: 0, y: -8 }}
              transition={{ duration: 0.2 }}
              className={styles.dropdown}
            >
              {labelInput.trim() !== "" &&
                !availableLabels.some(
                  (label) =>
                    label.toLowerCase() === labelInput.trim().toLowerCase()
                ) && (
                  <div
                    className={`${styles.dropdownItem} ${styles.newLabelItem}`}
                    onMouseDown={(e) => {
                      e.preventDefault();
                      handleAddLabel(labelInput.trim());
                    }}
                  >
                    <span style={{ fontSize: "16px" }}>+</span>
                    Create: <strong>{labelInput.trim()}</strong>
                  </div>
                )}
              {filteredLabels.map((label) => (
                <div
                  key={label}
                  className={styles.dropdownItem}
                  onMouseDown={(e) => {
                    e.preventDefault();
                    handleAddLabel(label);
                  }}
                >
                  {label}
                </div>
              ))}
            </motion.div>
          )}
        </AnimatePresence>

        {/* Error message */}
        {showError && error && (
          <div className={styles.errorMessage}>{error}</div>
        )}
      </div>
    );
  }
);

export default LabelsInput;