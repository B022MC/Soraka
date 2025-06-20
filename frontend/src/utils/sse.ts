export type SSECallback = (data: any) => void;

export class SSEClient {
  private source: EventSource | null = null;
  private reconnectTimer: number | null = null;
  private listeners: Record<string, SSECallback> = {};
  private heartbeatTimer: number | null = null;

  constructor(
    private url: string,
    private retry = 3000,
  ) {}

  private init() {
    this.source = new EventSource(this.url);
    // ignore broker heartbeat events
    this.source.addEventListener("ping", () => {});
    for (const [name, cb] of Object.entries(this.listeners)) {
      this.source.addEventListener(name, (e: MessageEvent) => {
        try {
          cb(JSON.parse(e.data));
        } catch {
          cb(e.data);
        }
      });
    }
    this.source.onerror = () => {
      if (this.reconnectTimer === null) {
        this.reconnectTimer = window.setTimeout(() => {
          this.reconnectTimer = null;
          this.source?.close();
          this.init();
        }, this.retry);
      }
    };
    this.source.onopen = () => {
      if (this.reconnectTimer) {
        window.clearTimeout(this.reconnectTimer);
        this.reconnectTimer = null;
      }
    };
  }

  connect(events: Record<string, SSECallback>) {
    this.listeners = events;
    this.init();
    if (this.heartbeatTimer === null) {
      this.heartbeatTimer = window.setInterval(() => {
        fetch("/heartbeat", { method: "POST" }).catch(() => {});
      }, 10000);
    }
  }

  close() {
    if (this.reconnectTimer) {
      window.clearTimeout(this.reconnectTimer);
      this.reconnectTimer = null;
    }
    if (this.heartbeatTimer) {
      window.clearInterval(this.heartbeatTimer);
      this.heartbeatTimer = null;
    }
    this.source?.close();
  }
}
