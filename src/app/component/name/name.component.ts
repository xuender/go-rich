import { Component, OnInit, Input } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { Obj } from '../../api/obj';
import { Config } from '../../api/config'

@Component({
  selector: 'rich-name',
  templateUrl: './name.component.html',
  styleUrls: ['./name.component.scss']
})
export class NameComponent implements OnInit {
  @Input()
  id: string
  name: string = ''
  del: boolean = false
  constructor(private http: HttpClient) { }

  ngOnInit() {
    console.log(this.id)
    let code
    switch (this.id[0]) {
      case 'I':
        code = 'items'
        break
      case 'C':
        code = 'customers'
        break
      default:
        console.error('ID错误', this.id)
    }
    this.http.get<Obj>(`${Config.URL}/api/${code}/${this.id}`)
      .subscribe(r => {
        this.name = r.name
        this.del = r.da && !r.da.startsWith('0001')
        if (this.del) {
          this.name = `[ ${r.name} ]`
        }
      })
  }
}
