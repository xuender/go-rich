import { Component } from '@angular/core';
import { ModalController } from '@ionic/angular';
import { FileUploader, FileItem, Headers } from 'ng2-file-upload';
import { Observable } from 'rxjs';

import { CustomerService } from '../customer.service';
import { URL } from '../../api/init'
import { SettingService } from '../../setting/setting.service';
import { Xlsx } from '../../setting/xlsxes/xlsx';
import { XlsxService } from '../../setting/xlsxes/xlsx.service';

@Component({
  selector: 'app-upload',
  templateUrl: './upload.page.html',
  styleUrls: ['./upload.page.scss'],
})
export class UploadPage {
  header: Headers = {
    name: 'xlsx',
    value: '',
  }
  headers: Headers[] = [this.header]
  uploader: FileUploader = new FileUploader({
    url: `${URL}/api/customers/file`,
    headers: this.headers,
    authToken: `Bearer ${localStorage.getItem("token")}`
  })
  select: Xlsx
  xlsxes$: Observable<Xlsx[]>
  constructor(
    public modalCtrl: ModalController,
    public customer: CustomerService,
    public setting: SettingService,
    private xlsxService: XlsxService,
  ) {
    this.xlsxes$ = this.xlsxService.all$
    this.uploader.onCompleteItem = (item: FileItem, r: string, status: number) => {
      if (status !== 200) {
        // TOOD 服务器错误
        return
      }
      this.modalCtrl.dismiss()
    }
  }

  async setup() {
    await this.setting.xlsxes()
    this.xlsxes$ = this.xlsxService.all$
  }

  cancel() {
    this.modalCtrl.dismiss();
  }
}
