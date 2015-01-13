os.File
=======

http://golang.org/pkg/os/

```
type File
    func Create(name string) (file *File, err error)
    func NewFile(fd uintptr, name string) *File
    func Open(name string) (file *File, err error)
    func OpenFile(name string, flag int, perm FileMode) (file *File, err error)
    func Pipe() (r *File, w *File, err error)
    func (f *File) Chdir() error
    func (f *File) Chmod(mode FileMode) error
    func (f *File) Chown(uid, gid int) error
    func (f *File) Close() error
    func (f *File) Fd() uintptr
    func (f *File) Name() string
    func (f *File) Read(b []byte) (n int, err error)
    func (f *File) ReadAt(b []byte, off int64) (n int, err error)
    func (f *File) Readdir(n int) (fi []FileInfo, err error)
    func (f *File) Readdirnames(n int) (names []string, err error)
    func (f *File) Seek(offset int64, whence int) (ret int64, err error)
    func (f *File) Stat() (fi FileInfo, err error)
    func (f *File) Sync() (err error)
    func (f *File) Truncate(size int64) error
    func (f *File) Write(b []byte) (n int, err error)
    func (f *File) WriteAt(b []byte, off int64) (n int, err error)
    func (f *File) WriteString(s string) (ret int, err error)
type FileInfo
    func Lstat(name string) (fi FileInfo, err error)
    func Stat(name string) (fi FileInfo, err error)
type FileMode
    func (m FileMode) IsDir() bool
    func (m FileMode) IsRegular() bool
    func (m FileMode) Perm() FileMode
    func (m FileMode) String() string
```
    
type File
---------

ファイルを表す。ファイルに対する操作、情報の取得を行う。

### ファイルを作る・開く ###

- Create
    - ファイル名を指定して新しくファイルを作る。
    - 既にファイルが存在する場合は空にしてから開く。
    - 読み込み・書き込み権限(O_RDWR)あり。
- NewFile
    - ファイルディスクリプタを指定してファイルを開く。
    - TODO: nameパラメータは何に使われる？Syncしたときにそこに保存されるわけではない。
- Open
    - ファイル名を指定してファイルを開く。
    - ファイルが存在しなければエラー(no such file or directory)。
    - 読み込み権限(O_RDONLY)のみ。
- OpenFile
    - flagとpermissionを指定してファイルを開く。

### ファイルに書き込む ###

- Write
    - []byteをファイルの現在の位置に書き込む。
- WriteAt
    - []byteをファイルの指定の位置に書き込む。
- WriteString
    - stringを現在の位置に書き込む。
- Truncate
    - ファイルのサイズを変更する。

### ファイルのメタ情報を編集する ###

- Chmod
- Chown
- Chtimes
- Sync
    - Fileの内容を実際のファイルに反映する。
    - Closeするときに呼ばれる？

### その他 ###

- Chdir
    - カレントディレクトリをファイルのパス変更。
    - syscall.Fchdirが呼ばれてる。

type FileInfo
-------------

ファイルの情報の入れ物。
os.Statとos.Lstatの戻り値に使われている。

type FileMode
-------------

ファイルのモードを表す。これもstatコマンドで取得できるやつ。
それぞれのモードに対応する定数がFileMode型で定義されており、
それを取得するメソッドが生えている。
