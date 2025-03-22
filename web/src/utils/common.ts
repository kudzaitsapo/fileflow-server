export function formatBytes(bytes: number, decimals = 2) {
  if (!+bytes) return "0 Bytes";

  const k = 1024;
  const dm = decimals < 0 ? 0 : decimals;
  const sizes = ["Bytes", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"];

  const i = Math.floor(Math.log(bytes) / Math.log(k));

  return `${parseFloat((bytes / Math.pow(k, i)).toFixed(dm))} ${sizes[i]}`;
}

export function formatDateTime(input: string) {
  // Convert input to Date object (handles both Date objects and date strings)
  const date = new Date(input);

  // Check if the date is valid
  if (isNaN(date.getTime())) {
    return "Invalid date input";
  }

  // Array of month abbreviations
  const months = [
    "Jan",
    "Feb",
    "Mar",
    "Apr",
    "May",
    "Jun",
    "Jul",
    "Aug",
    "Sep",
    "Oct",
    "Nov",
    "Dec",
  ];

  // Get date components
  const month = months[date.getMonth()];
  const day = date.getDate();
  const year = date.getFullYear();

  // Get time components and pad with leading zeros if needed
  const hours = date.getHours().toString().padStart(2, "0");
  const minutes = date.getMinutes().toString().padStart(2, "0");

  // Construct the formatted string
  return `${month} ${day}, ${year} at ${hours}:${minutes}hrs`;
}
