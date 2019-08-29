declare var external;

type RendererCallback = (event: string, value: any) => void;

export class IPCRenderer {
  constructor(private eventTable: { [key: string]: RendererCallback } = {}) {}

  public send(event: string, value: any) {
    external.invoke(`${event}:${JSON.stringify(value)}`);
  }

  public on(event: string, callback: RendererCallback) {
    this.eventTable[event] = callback;
  }

  public trigger(event: string, value: string) {
    const cb = this.eventTable[event];
    if (cb !== undefined) {
      cb(event, JSON.parse(value));
    } else {
      console.log("No event handler");
    }
  }
}
