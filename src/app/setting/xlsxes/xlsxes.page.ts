import { Component } from '@angular/core';
import { ModalController, ActionSheetController } from '@ionic/angular';
import { Observable } from 'rxjs';

import { ObjsPage } from '../../api/objs.page';
import { Xlsx } from './xlsx';
import { XlsxService } from './xlsx.service';
import { XlsxPage } from './xlsx/xlsx.page';

@Component({
  selector: 'app-xlsxes',
  templateUrl: './xlsxes.page.html',
  styleUrls: ['./xlsxes.page.scss'],
})
export class XlsxesPage extends ObjsPage<Xlsx> {
  xlsxes$: Observable<Xlsx[]>
  constructor(
    public xlsxService: XlsxService,
    protected modalCtrl: ModalController,
    protected actionSheetCtrl: ActionSheetController,
  ) {
    super(modalCtrl, actionSheetCtrl)
    this.xlsxes$ = this.xlsxService.all$
  }
  get service() { return this.xlsxService }
  get page() { return XlsxPage }
  get newObj() { return {} }
  get title() { return 'Excel格式' }
}
