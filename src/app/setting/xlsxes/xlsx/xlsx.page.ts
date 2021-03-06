import { Component } from '@angular/core';
import { NavParams, ModalController } from '@ionic/angular';
import { invert } from 'lodash'

import { Xlsx } from '../xlsx';
import { Observable } from 'rxjs';
import { Ext } from '../../../ext/ext';
import { XlsxService } from '../xlsx.service';
import { ObjPage } from '../../../api/obj.page';
import { ExtService } from '../../../ext/ext.service';

@Component({
  selector: 'app-xlsx',
  templateUrl: './xlsx.page.html',
  styleUrls: ['./xlsx.page.scss'],
})
export class XlsxPage extends ObjPage<Xlsx>{
  abc = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'.split('')
  types = ['客户', '商品', '交易']
  extsByCustomer$: Observable<Ext[]>
  extsByItem$: Observable<Ext[]>
  constructor(
    public extService: ExtService,
    private xlsxService: XlsxService,
    protected modalCtrl: ModalController,
    protected navParams: NavParams
  ) {
    super(modalCtrl, navParams)
    this.extsByCustomer$ = this.extService.extsByCustomer$
    this.extsByItem$ = this.extService.extsByItem$
    this.obj.map = invert(this.obj.map)
    for (const k in this.obj.map) {
      if (this.obj.map[k]) {
        this.obj.map[k] = parseInt(this.obj.map[k]) + 1
      }
    }
  }

  get service() { return this.xlsxService }

  async save() {
    console.log('x save')
    for (const k in this.obj.map) {
      if (!this.obj.map[k]) {
        delete this.obj.map[k]
      } else {
        this.obj.map[k] = this.obj.map[k] - 1
      }
    }
    this.obj.map = invert(this.obj.map)
    await super.save()
  }

  async customer() {
    await this.extService.extCustomer()
    this.extsByCustomer$ = this.extService.extsByCustomer$
  }

  async item() {
    await this.extService.extItem()
    this.extsByItem$ = this.extService.extsByItem$
  }
}
