import { Component, ViewChild, ElementRef } from '@angular/core';
import { ModalController } from '@ionic/angular';
import { Observable } from 'rxjs';

import { Xlsx } from '../../setting/xlsxes/xlsx';
import { CustomerService } from '../customer.service';
import { SettingService } from '../../setting/setting.service';
import { XlsxService } from '../../setting/xlsxes/xlsx.service';

@Component({
  selector: 'app-download',
  templateUrl: './download.page.html',
  styleUrls: ['./download.page.scss'],
})
export class DownloadPage {
  @ViewChild('downloadZipLink')
  private downloadZipLink: ElementRef;
  xlsxes$: Observable<Xlsx[]>
  params = { excel: '' }
  constructor(
    private customerService: CustomerService,
    public modalCtrl: ModalController,
    public customer: CustomerService,
    public setting: SettingService,
    private xlsxService: XlsxService,
  ) {
    this.xlsxes$ = this.xlsxService.all$
  }

  async setup() {
    await this.setting.xlsxes()
    this.xlsxes$ = this.xlsxService.all$
  }

  cancel() {
    this.modalCtrl.dismiss();
  }

  async download() {
    const blob = await this.customerService.download(this.params).toPromise()
    const url = window.URL.createObjectURL(blob);
    const link = this.downloadZipLink.nativeElement;
    link.href = url;
    link.download = 'customers.xlsx';
    link.click();
    window.URL.revokeObjectURL(url);
  }
}
