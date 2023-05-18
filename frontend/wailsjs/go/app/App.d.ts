// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {app} from '../models';
import {config} from '../models';
import {starrs} from '../models';

export function Ask(arg1:string,arg2:string):Promise<boolean>;

export function CheckUpdate():Promise<app.Release>;

export function CreateShortcut():Promise<string>;

export function DownloadUpdate():Promise<app.UpdateInfo>;

export function ErrorDialog(arg1:string,arg2:string):Promise<void>;

export function GetConfig():Promise<config.Settings>;

export function Languages():Promise<{[key: string]: string}>;

export function LaunchInstaller(arg1:string):Promise<string>;

export function OpenFolder(arg1:string):Promise<string>;

export function PickFile(arg1:string,arg2:string,arg3:string):Promise<string>;

export function PickFolder(arg1:string):Promise<string>;

export function Quit():Promise<void>;

export function RemoveInstance(arg1:number,arg2:string):Promise<app.SavedInstance>;

export function SaveConfigItem(arg1:string,arg2:any,arg3:boolean):Promise<app.ConfigSaved>;

export function SaveInstance(arg1:number,arg2:starrs.AppConfig,arg3:boolean):Promise<app.SavedInstance>;

export function Version():Promise<app.Version>;
