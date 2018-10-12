import { HttpClient } from '@angular/common/http';
import { Pipe, PipeTransform } from '@angular/core';

import { Config } from '../api/config'
@Pipe({
  name: 'toName',
  pure: false
})
export class ToNamePipe implements PipeTransform {
  static cache: any = {}
  constructor(private http: HttpClient) { }
  transform(id: string): any {
    if (!id) { return '' }
    if (!(id in ToNamePipe.cache)) {
      let name
      switch (id[0]) {
        case 'I':
          name = 'items'
          break
        case 'C':
          name = 'customers'
          break
        default:
          console.error('ID错误', id)
      }
      this.http.get(`${Config.URL}/api/${name}/${id}`)
        .subscribe((r: any) => ToNamePipe.cache[id] = r.name)
    }
    return ToNamePipe.cache[id];
  }

}
