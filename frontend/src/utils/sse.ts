export type SSECallback = (data: any) => void;

export class SSEClient {
  private source: EventSource | null = null;
  constructor(private url: string) {}

  connect(events: Record<string, SSECallback>) {
    this.source = new EventSource(this.url);
    for (const [name, cb] of Object.entries(events)) {
      this.source.addEventListener(name, (e: MessageEvent) => {
        try {
          cb(JSON.parse(e.data));
        } catch {
          cb(e.data);
        }
      });
    }
  }

  close() {
    this.source?.close();
  }
}
