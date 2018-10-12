import { environment } from "../../environments/environment";

export class Config {
  public static get URL() {
    if (environment.production) {
      return '';
    }
    return 'http://localhost:6181';
  }
}
