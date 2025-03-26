/**
 * Adapter Design Pattern
 *
 * Intent: Provides a unified interface that allows objects with incompatible
 * interfaces to collaborate.
 */

/**
 * The Target defines the domain-specific interface used by the client code.
 */
export interface FileViewer {
    openFile(file: string): string 
}

/**
 * This is the code I WANT to use in the client, but cant (directly)
 * The Adaptee contains some useful behavior, but its interface is incompatible
 * with the existing client code. The Adaptee needs some adaptation before the
 * client code can use it.
 * (3rd party library youw ant to use, but is just incompatible with client)
 * Instead of making many modifications to client, we can just point the client to use an adapter,
 * have the adapter wrap the service, and the client doesnt have to change anything
 */
class videoConversion {
    public convertWebmToMp4(file: string): string {
        const fileExtension = this.getVideoFileExtension(file);
        if (fileExtension === FileTypes.Mp4) {
            return file;
        }

        const mp4Extension = FileTypes.Mp4;
        return file.replace(fileExtension, mp4Extension);
    }

    private getVideoFileExtension(file: string): string {
        const splitFile = file.split('.');
        return splitFile[splitFile.length -1]
    }
}

export enum FileTypes {
    Jpg = "jpg",
    Png = "png",
    Mp3 = "mp3",
    Mp4 = "mp4",
    Webm = "webm"
}

/**
 * The Adapter makes the Adaptee's interface compatible with the Target's
 * interface.
 */
class FileViewerAdapter implements FileViewer {
    constructor(private vc: videoConversion) {}

    public openFile(file: string): string {
        // Determine file type
        const fileExteion = this.getFileExtension(file);
        // Wrapped the service into the adapter, now we can call the previously incompatible method from our adapter
        // Since the adapter implements the Client interface. client can call all methods on the interface with confidence
        if (fileExteion === FileTypes.Webm) {
            // convert webM to mp4
            return this.vc.convertWebmToMp4(file);
        }

        // Picture
        return this.convertToPng(file);
    }

    private convertToPng(file: string): string {
        const fileExtension = this.getFileExtension(file);
        const pngExtension = FileTypes.Png;
        return file.replace(fileExtension, pngExtension);
    }

    private getFileExtension(file: string): string {
        const splitFile = file.split('.');
        return splitFile[splitFile.length -1]
    }
}


/**
 * The client code works with objects that implements the TaxiCalculator
 * interface, so we can use the adapter to reuse the incompatible library
 */
function client(fileViewer: FileViewerAdapter): void {
    console.log("Want to open a jpg; SW only supports png.  Will implicitly convert on open");
    const myFile = "cats.jpg";
    console.log(`Original file: ${myFile}`);
    const file = fileViewer.openFile(myFile);
    console.log(`The file has been converted before opening: ${file}`);

    // Call methods available only in 3rd party service
    console.log("Want to open video file. I want use non compatible service THROUGH the adapter");
    const originalVideo = 'replay.webm';
    console.log(`Original video: ${originalVideo}`);
    const videoFile = fileViewer.openFile(originalVideo);
    console.log(`The video file has been converted before opening: ${videoFile}`);
}

const incompatibleLibrary = new videoConversion();
const adaptedLibrary = new FileViewerAdapter(incompatibleLibrary);
client(adaptedLibrary);


// Its paramount that the Adapter implements a client interface, and the client has to call the interface.
// This ensures decoupled code.
// The reason why client cant call directly to the Adaptee/ 3rd party is because of incompatible.  
// Client code doesnt need to change, but a call to the adapter passing the Adaptee is a must. The return object is
// just an instance of the client interface, which is what the client expects / knows how to work with.
