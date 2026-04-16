/**
 * Format string as currency IDR
 * @param value number to format
 */
export const formatCurrency = (value: number): string => {
  return new Intl.NumberFormat('id-ID', {
    style: 'currency',
    currency: 'IDR',
  }).format(value)
}

/**
 * Capitalize first letter of each word
 * @param str text to capitalize
 */
export const capitalizeWords = (str: string): string => {
  return str.replace(/\b\w/g, (char) => char.toUpperCase())
}
