export interface AppState {
  system: SystemSettings;
  client: ClientStatus;
  lcu: LcuStatus;
}

export interface SystemSettings {
  theme: string;       // light / dark
  sysTime: string;     // 当前系统时间
}

export interface ClientStatus {
  clientPath: string;  // 客户端路径
}

export interface LcuStatus {
  online: boolean;     // 是否在线
  port: string;        // 端口
  token: string;       // 凭证
}
