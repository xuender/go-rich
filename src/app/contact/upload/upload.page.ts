import { Component, OnInit } from '@angular/core';
import { CustomerService } from '../../api/customer.service';
import { ModalController } from '@ionic/angular';
import { FileUploader, FileItem, Headers } from 'ng2-file-upload';
import { HttpClient } from '@angular/common/http';
import { URL } from '../../api/init'
import { SettingService } from '../../api/setting.service';
import { Xlsx } from '../../api/xlsx';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-upload',
  templateUrl: './upload.page.html',
  styleUrls: ['./upload.page.scss'],
})
export class UploadPage implements OnInit {
  header: Headers = {
    name: 'xlsx',
    value: '',
  }
  headers: Headers[] = [this.header]
  uploader: FileUploader = new FileUploader({ url: `${URL}/api/customers/file`, headers: this.headers });
  select: Xlsx
  xlsxes: Observable<Xlsx[]>
  constructor(
    private http: HttpClient,
    public modalCtrl: ModalController,
    public customer: CustomerService,
    public setting: SettingService,
  ) {
    this.xlsxes = this.http.get<Xlsx[]>(`${URL}/api/xlsxes`)
    this.uploader.onCompleteItem = (item: FileItem, r: string, status: number) => {
      if (status !== 200) {
        // TOOD 服务器错误
        return
      }
      this.modalCtrl.dismiss()
    }
  }

  ngOnInit() {
  }
  cancel() {
    this.modalCtrl.dismiss();
  }
}
