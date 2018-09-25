import { CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { DownloadPage } from './download.page';

describe('UploadPage', () => {
  let component: DownloadPage;
  let fixture: ComponentFixture<DownloadPage>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ DownloadPage ],
      schemas: [CUSTOM_ELEMENTS_SCHEMA],
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(DownloadPage);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
