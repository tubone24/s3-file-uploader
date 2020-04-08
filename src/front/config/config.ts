import configJson from '@/config/config.json'

export type Config = typeof configJson;
export type FileOption = typeof configJson.fileOption;
export const config = configJson;
export const fileOption = configJson.fileOption;
